generate_users:
	@echo "Generating users.proto"
	./generate.sh user ../server/users/src/proto-gen/ ../client/packages/users_repository/lib/src/gen/proto/

generate_auth:
	@echo "Generating auth.proto"
	./generate.sh auth ../server/auth/src/proto-gen/

all: generate_users generate_auth
