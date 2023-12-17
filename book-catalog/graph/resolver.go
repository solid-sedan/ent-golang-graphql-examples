package graph

import (
	"book-catalog/controllers"
	"book-catalog/ent"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/zerolog"
)

// Resolver is the resolver root.
type Resolver struct {
	BookClient controllers.BookController
}

// NewSchema creates a graphql executable schema.
func NewSchema(logger zerolog.Logger, client *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			BookClient: controllers.NewBookController(logger, client),
		},
	})
}
