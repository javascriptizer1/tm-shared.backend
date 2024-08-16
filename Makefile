lint:
	docker run --rm -v .:/src -w /src golangci/golangci-lint:v1.59.1 golangci-lint --config .golangci.pipeline.yaml run
