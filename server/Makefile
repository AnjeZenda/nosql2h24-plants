
gen-proto:
	protoc -I api .\api\plantsapi\plants_api.proto --go_out=./internal/pb --go_opt=paths=source_relative --go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative

proto:
	if not exist ".\internal\pb\plantsapi" mkdir ".\internal\pb\plantsapi"
	protoc -I api \
	       -I$(GOPATH)/pkg/mod \
	       -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
	       -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0 \
	       --grpc-gateway_out=logtostderr=true:./internal/pb/plantsapi \
	       --swagger_out=allow_merge=true,merge_file_name=api:. \
	       --go_out=plugins=grpc:./internal/pb/plantsapi .\api\plantsapi\plants_api.proto
