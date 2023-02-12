.PHONY: run
run:
	go run main.go

.PHONY: docs
docs:
	swag init


.PHONY: postgres-dev
postgres-dev:
	docker run --name rent-cat-postgres \
		-e POSTGRES_DB=rent-cat-db \
		-e POSTGRES_USER=rent-cat-user \
		-e POSTGRES_PASSWORD=rent-cat-pass \
		-d postgres
