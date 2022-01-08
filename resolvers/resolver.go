package resolvers

import (
	ent "170-ag/ent/generated"
	"170-ag/resolvers/generated"
	"github.com/99designs/gqlgen/graphql"
)

type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
