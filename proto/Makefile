# RELATIVE PATH SHOULD DEFINITELY BE MODIFIED

all: users rides 

# users ms 

users_server:
	@echo "Generating users.proto for server"
	./gen_server.sh users ../server/users/src/proto-gen/

users_client:
	@echo "Generating users.proto for client"
	./gen_client.sh users ../client/packages/users_repository/lib/src/gen/proto/

users: generate_users_server generate_users_client
=======
generate_users_server:
	@echo "Generating users.proto for client"
	./generate_users_server.sh user ../server/users/src/proto-gen/

generate_users_client:
	@echo "Generating users.proto for client"
	./generate_users_client.sh user ../client/packages/users_repository/lib/src/gen/proto/

all: generate_users_server generate_users_client
