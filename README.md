# OpenTelemetry Example

## Objective

Write an example on how to use OpenTelemetry with:
1. HTTP Client
2. gRPC Server
3. Postgres Client
4. Redis Client

## Running

Start Jaeger:

```shell
make deps
```

Run the application:

```shell
make run
```

Call the API:
```shell
grpcurl \
  -d '{ "cat": { "name": "Cat1", "id": "cat1" } }' \
  -plaintext localhost:8080 \
  api.v1.CatService/CreateCat
```

```shell
grpcurl \
  -plaintext localhost:8080 \
  api.v1.CatService/ListCats
```

Find the traces in [Jaeger UI](http://localhost:16686).