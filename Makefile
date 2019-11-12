.PHONY: run test

run:
	go run cmd/app/main.go

test:
	cd pkg && go test ./... && cd ..