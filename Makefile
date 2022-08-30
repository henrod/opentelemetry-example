.PHONY: run
run:
	go run ./...

.PHONY: deps
deps:
	docker compose up -d

.PHONY: protos
protos:
	cd proto/cat && buf push
	buf generate buf.build/henrod/cat