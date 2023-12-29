build_proto:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/price_checker.proto

run:
	go build -o app && ./app