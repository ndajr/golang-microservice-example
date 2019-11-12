.PHONY: run test

run:
	GIT_COMMIT=$$(git rev-parse --short=10 HEAD) go run cmd/app/main.go

test:
	cd pkg && go test ./... && cd ..