package graphql

import (
	"book-catalog/controllers"
	"book-catalog/ent"
	"book-catalog/graphql/generated"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/zerolog"
)

// Resolver is the resolver root.
type Resolver struct {
	BookClient controllers.BookController
}

// NewSchema creates a graphql executable schema.
func NewSchema(logger zerolog.Logger, client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			BookClient: controllers.NewBookController(logger, client),
		},
	})
}
