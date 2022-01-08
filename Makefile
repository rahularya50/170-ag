.PHONY: rebuild-graphql ent

rebuild-graphql:
	go run github.com/99designs/gqlgen generate

ent:
	go generate ./...

serve-site:
	go run ./cmd/site
