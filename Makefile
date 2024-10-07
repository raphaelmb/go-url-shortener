run:
	@go build -o bin/url_shortener ./cmd/api && ./bin/url_shortener

.PHONY: run