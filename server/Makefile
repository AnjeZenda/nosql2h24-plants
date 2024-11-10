gen-proto:
	if not exist ".\internal\pb\plantsapi" mkdir ".\internal\pb\plantsapi"
	protoc -I api \
	       -I$(GOPATH)/pkg/mod \
	       -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
	       -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0 \
		   --go_out=internal/pkg/pb --go_opt=paths=source_relative \
	       --grpc-gateway_out=internal/pkg/pb --grpc-gateway_opt=paths=source_relative \
	       --go-grpc_out=internal/pkg/pb --go-grpc_opt=paths=source_relative .\api\plants\v1\plants.proto


proto:
	rm -rf internal/pkg/pb && mkdir -p internal/pkg/pb
	protoc -I api -I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		--go_out=internal/pkg/pb --go_opt=paths=source_relative \
		--go-grpc_out=internal/pkg/pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=internal/pkg/pb --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out gen/docs \
		./api/**/**/*.proto
