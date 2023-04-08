
wire: ## Generate wire_gen.go
	cd pkg/di && wire


run: ## Start application
	go run ./cmd/api