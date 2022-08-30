package v1

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"opentelemetry-example/domain/entities"
	v1 "opentelemetry-example/protogen/go/api/v1"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"
)

type CatService struct {
	httpClient     *http.Client
	storageGateway StorageGateway
}

func NewCatService(
	httpClient *http.Client,
	storageClient StorageGateway,
) (*CatService, error) {
	return &CatService{
		httpClient:     httpClient,
		storageGateway: storageClient,
	}, nil
}

func (catService *CatService) CreateCat(ctx context.Context, request *v1.CreateCatRequest) (*v1.CreateCatResponse, error) {
	cat := &entities.Cat{
		ID:   request.Cat.Id,
		Name: request.Cat.Name,
	}

	err := catService.storageGateway.CreateCat(ctx, cat)
	if err != nil {
		log.Printf("failed to create cat in gateway: %s", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &v1.CreateCatResponse{
		Cat: request.Cat,
	}, nil
}

func (catService *CatService) ListCats(ctx context.Context, _ *v1.ListCatsRequest) (*v1.ListCatsResponse, error) {
	cats, err := catService.storageGateway.ListCats(ctx)
	if err != nil {
		log.Printf("failed to list cats in gateway: %s", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	pCats := make([]*v1.Cat, 0, len(cats))
	for _, cat := range cats {
		pCats = append(pCats, &v1.Cat{
			Id:   cat.ID,
			Name: cat.Name,
		})
	}

	return &v1.ListCatsResponse{
		Cats: pCats,
	}, nil
}

func (catService *CatService) GetFact(ctx context.Context, _ *v1.GetFactRequest) (*v1.GetFactResponse, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://catfact.ninja/fact", nil)
	if err != nil {
		log.Printf("failed to build request: %s", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp, err := catService.httpClient.Do(request)
	if err != nil {
		log.Printf("failed to get catfact url: %s", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	defer func() { _ = resp.Body.Close() }()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body: %s", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	catFact := new(entities.CatFact)
	err = json.Unmarshal(respBytes, catFact)
	if err != nil {
		log.Printf("failed to unmarshal cat fact: %s", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &v1.GetFactResponse{
		Fact: catFact.Fact,
	}, nil
}
