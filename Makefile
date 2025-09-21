.PHONY: proto

run:
	go run cmd/server/main.go

proto:
	protoc --go_out=. --go-grpc_out=. proto/server/v1/user.proto