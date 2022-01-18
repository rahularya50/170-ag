.PHONY: ent graphql proto relay flow view

# generate Go interfaces to interact with data models from Ent schema
ent:
	go generate ./...

# autogenerate (possibly stubbed) resolvers for Ents and oher models in graphQL schema
graphql:
	go run github.com/99designs/gqlgen generate

proto:
	protoc -I proto/ --go_out . proto/*.proto

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
