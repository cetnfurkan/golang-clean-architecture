build: ent grpc swagger

ent:
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/schemaconfig ./entity --target=target/ent

grpc:
	protoc --go_out=./ --go-grpc_out=./ --grpc-gateway_out=./ --grpc-gateway_opt=generate_unbound_methods=true --grpc-gateway_opt=grpc_api_configuration=./gw_mapping.yaml  --proto_path=./protos/golang-clean-architecture user.proto && \
	protoc --openapiv2_out=./target/grpc/user --openapiv2_opt=grpc_api_configuration=./gw_mapping.yaml --proto_path=./protos/golang-clean-architecture user.proto

swagger:
	go run github.com/swaggo/swag/cmd/swag@v1.8.12 init -o target/swagger

run:
	go build && ./golang-clean-architecture