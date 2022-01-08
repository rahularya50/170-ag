.PHONY: ent graphql relay flow view

# generate Go interfaces to interact with data models from Ent schema
ent:
	go generate ./...

# autogenerate (possibly stubbed) resolvers for Ents and oher models in graphQL schema
graphql:
	go run github.com/99designs/gqlgen generate

# compile graphQL fragments in frontend and generate Flow typings
relay:
	cd frontend && yarn relay

# verify all Flow typings
flow:
	cd frontend && yarn flow

# serve backend (run `yarn start` for frontend)
view:
	go run ./cmd/site
