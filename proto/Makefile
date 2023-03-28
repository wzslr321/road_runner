generate_users_server:
	@echo "Generating users.proto for client"
	./generate_users_server.sh user ../server/users/src/proto-gen/

generate_users_client:
	@echo "Generating users.proto for client"
	./generate_users_client.sh user ../client/packages/users_repository/lib/src/gen/proto/

all: generate_users_server generate_users_client

client: generate_users_client

server: generate_users_server

ci_client: :
	@echo "Generating users.proto for client"
	./generate_users_client_ci.sh user ../client/packages/users_repository/lib/src/gen/proto/
