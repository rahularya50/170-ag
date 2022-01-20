.PHONY: ent graphql proto relay flow view connect-db enter-db

# generate Go interfaces to interact with data models from Ent schema
ent:
	go generate ./...

# autogenerate (possibly stubbed) resolvers for Ents and oher models in graphQL schema
graphql:
	go run github.com/99designs/gqlgen generate

proto:
	protoc -I proto/ --go_out . proto/*.proto --go-grpc_out .

# compile graphQL fragments in frontend and generate Flow typings
relay:
	cd frontend && yarn relay

# verify all Flow typings
flow:
	cd frontend && yarn flow

# serve backend (run `yarn start` for frontend)
view: export SESSION_KEY = abcdef
view: export ENV = dev

view:
	go run ./cmd/site

connect-db:
	./cloud_sql_proxy -enable_iam_login -instances=formidable-gate-337712:us-west2:cs170-db=tcp:5432

enter-db:
	psql --host=localhost --port=5432 --user=rahularya50@gmail.com --dbname=autograder
