START_LOG = @echo "================================================= START OF LOG ==================================================="
END_LOG = @echo "================================================== END OF LOG ===================================================="

.PHONY: env
env: ./.env.develop.tmpl
	$(START_LOG)
	@cp ./.env.develop.tmpl ./.env
	@echo "Environment file created at ./.env"
	$(END_LOG)
	
.PHONY: run
run:
	$(START_LOG)
	@docker compose \
		-f ./deployments/compose.packages.yaml up \
		--build -d
	$(END_LOG)

.PHONY: swagger
swagger:
	$(START_LOG)
	@docker run --rm -v $(shell pwd):/code ghcr.io/swaggo/swag:latest i -g ./cmd/user_microservice/main.go -o ./api
	@go mod tidy
	$(END_LOG)