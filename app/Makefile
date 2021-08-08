.phony: generate install build test go

install:
	go get -x google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		github.com/envoyproxy/protoc-gen-validate \
		github.com/gogo/protobuf

generate:
	protoc --proto_path=api/proto/v1 --proto_path=./third_party \
		--go_out=./pkg/api/v1 --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/api/v1 --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pkg/api/v1 --grpc-gateway_opt paths=source_relative --grpc-gateway_opt logtostderr=true \
		--validate_out=lang=go,paths=source_relative:./pkg/api/v1 \
		--openapiv2_out ./api/openapiv2/v1 --openapiv2_opt logtostderr=true,allow_merge=true,merge_file_name=listing.json \
		listing.proto
	cp ./api/openapiv2/v1/listing.swagger.json ./openapiv2/

go: 
	go mod tidy

buildgw:
	go build -o ./out/gateway -mod vendor cmd/gateway/main.go
	
buildlisting:	
	go build -o ./out/listing -mod vendor cmd/listing/main.go
