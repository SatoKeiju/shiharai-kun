.PHONY: goa-gen
goa-gen:
	go run goa.design/goa/v3/cmd/goa@v3.19.1 gen github.com/SatoKeiju/shiharai-kun/design;

.PHONY: lint
lint: ## lint実行
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run
