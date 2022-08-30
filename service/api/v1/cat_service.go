package v1

import (
	"context"
	"encoding/json"
	"fmt"
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
	var err error
	cat := &entities.Cat{
		ID:   request.Cat.GetId(),
		Name: request.Cat.GetName(),
		Fact: "",
	}

	cat.Fact, err = catService.getFact(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get cat fact: %w", err)
	}

	err = catService.storageGateway.CreateCat(ctx, cat)
	if err != nil {
		log.Printf("failed to create cat in gateway: %s", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	pCat := &v1.Cat{
		Id:   cat.ID,
		Name: cat.Name,
		Fact: cat.Fact,
	}

	return &v1.CreateCatResponse{
		Cat: pCat,
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
			Fact: cat.Fact,
		})
	}

	return &v1.ListCatsResponse{
		Cats: pCats,
	}, nil
}

type CatFact struct {
	Fact string
}

func (catService *CatService) getFact(ctx context.Context) (string, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://catfact.ninja/fact", nil)
	if err != nil {
		return "", fmt.Errorf("failed to build request: %w", err)
	}

	resp, err := catService.httpClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("failed to get catfact url: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	catFact := new(CatFact)
	err = json.Unmarshal(respBytes, catFact)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal cat fact: %w", err)
	}

	return catFact.Fact, nil
}
