
gen-proto:
	protoc -I api .\api\plantsapi\plants_api.proto --go_out=./internal/pb --go_opt=paths=source_relative --go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative

proto:
	rm -rf internal/pkg/pb && mkdir -p internal/pkg/pb
	protoc -I api -I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		--go_out=internal/pkg/pb --go_opt=paths=source_relative \
		--go-grpc_out=internal/pkg/pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=internal/pkg/pb --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out gen/docs \
		./api/**/**/*.proto
