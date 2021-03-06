.PHONY: grpc grpc_clean

generate_grpc: grpc_clean
	protoc --proto_path rpc/ --proto_path /Users/admin/go/src/github.com/golang/protobuf/ptypes/timestamp --go_out=plugins=grpc:rpc rpc/account.proto
	protoc --proto_path rpc/ --proto_path /Users/admin/go/src/github.com/golang/protobuf/ptypes/timestamp --go_out=plugins=grpc:rpc rpc/order.proto

grpc_clean:
	rm rpc/*.go || true

remove_omitempty:
	ls rpc/*.pb.go | xargs -n1 -IX bash -c "sed -e 's/,omitempty//' X > X.tmp && mv X{.tmp,}"

grpc: grpc_clean generate_grpc remove_omitempty

migrate:
	DB_NAME=crud tern migrate --migrations ./migrations

migrate_undo:
	DB_NAME=crud tern migrate --migrations ./migrations --destination -1

migrate_test:
	DB_NAME=crud_test tern migrate --migrations ./migrations

migrate_test_undo:
	DB_NAME=crud_test tern migrate --migrations ./migrations --destination -1

local_dev:
	docker-compose up db

connect_db:
	docker-compose exec db psql -U postgres -d crud