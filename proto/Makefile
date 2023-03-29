# RELATIVE PATH SHOULD DEFINITELY BE MODIFIED

all: users rides companies

# users ms 

users_server:
	@echo "Generating users.proto for server"
	./gen_server.sh users ../server/users/src/proto-gen/

users_client:
	@echo "Generating users.proto for client"
	./gen_client.sh users ../client/packages/users_repository/lib/src/gen/proto/

users: users_client users_server 

ci_client:
	@echo "Generating users.proto for client"
	./gen_client_ci.sh users ../client/packages/users_repository/lib/src/gen/proto/

# rides ms

rides_server:
	@echo "Generating rides.proto for server"
	./gen_server.sh rides ../server/rides/src/proto-gen/

rides_client:
	@echo "Generating rides.proto for client"
	./gen_client.sh rides ../client/packages/rides_repository/lib/src/gen/proto/

rides: rides_client rides_server 

# companies ms

companies_server:
	@echo "Generating companies.proto for server"
	./gen_server.sh companies ../server/companies/src/proto-gen/

companies_client:
	@echo "Generating companies.proto for client" 
	./gen_client.sh companies ../client/packages/companies_repository/lib/src/gen/proto/

companies: companies_server companies_client
