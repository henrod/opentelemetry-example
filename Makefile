.PHONY: run
run:
	go run ./...

.PHONY: deps
deps:
	docker compose up -d