# OpenTelemetry Example

## Objective

Write an example on how to use OpenTelemetry with:
1. HTTP Client
2. HTTP Server
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

Find the traces in [Jaeger UI](localhost:16686).