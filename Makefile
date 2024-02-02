protoc:
	protoc \
      --go_out=proto/impl \
      --go_opt=paths=import \
      --go-grpc_out=proto/impl \
      --go-grpc_opt=paths=import \
      proto/source/**/*.proto
