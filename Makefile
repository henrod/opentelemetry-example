.PHONY: run
run:
	go run ./...

.PHONY: deps
deps:
	docker compose up -d

.PHONY: migrations
migrations:
	@# This is not the best way to migrate, improve it later.
	psql postgresql://postgres:@localhost:5432/postgres\?sslmode=disable -f gateway/postgres/migrations/migrations.sql

.PHONY: protos
protos:
	cd proto/cat && buf push
	buf generate buf.build/henrod/cat