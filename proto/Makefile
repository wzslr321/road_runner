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

client: generate_users_client

server: generate_users_server

ci_client:
	@echo "Generating users.proto for client"
	./gen_client_ci.sh user ../client/packages/users_repository/lib/src/gen/proto/

# rides ms

rides_server:
	@echo "Generating rides.proto for server"
	./gen_server.sh rides ../server/rides/src/proto-gen/

rides_client:
	@echo "Generating rides.proto for client"
	./gen_client.sh rides ../client/packages/rides_repository/lib/src/gen/proto/

rides: generate_rides_server generate_rides_client
