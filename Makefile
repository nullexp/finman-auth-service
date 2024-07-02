.PHONY: dev-run install buf lint test docker-build docker-run

install:
	@go mod tidy
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

buf:
	mkdir -p "./proto/user/v1"
	curl -o ./proto/user/v1/user.proto https://raw.githubusercontent.com/nullexp/finman-user-service/main/proto/user/v1/user.proto
	@env PATH="$$PATH:$$(go env GOPATH)/bin" buf generate --template proto/buf.gen.yaml proto
	@echo "✅ buf done!"
	rm -rf "./proto/user"

buf-win:
	mkdir ".\proto\user\v1"
	curl -o .\proto\user\v1\user.proto https://raw.githubusercontent.com/nullexp/finman-user-service/main/proto/user/v1/user.proto
	@set PATH=%PATH%;%GOPATH%\bin
	@buf generate --template proto\buf.gen.yaml proto
	@echo "✅ buf done!"
	rmdir /S /Q ".\proto\user"

run:
	go run ./cmd

lint:
	gofumpt -l -w .
	golangci-lint run -v

test:
	go test ./...

docker-build:
	docker build -t finman-auth-service .

docker-run:
	docker run -p 8080:8080 finman-auth-service

docker-compose-up:
	docker-compose up --build 

docker-compose-down:
	docker-compose down --volumes