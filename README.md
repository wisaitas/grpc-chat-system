# install go plugin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# install go dependencies
go get google.golang.org/genproto/googleapis/api/annotations
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

# install protobuf compliler
## linux
brew install protobuf

## window
choco install protobuf