package graph

import (
	"book-catalog/controllers"
	"book-catalog/ent"

	"github.com/99designs/gqlgen/graphql"
	"github.com/atlas-health/atlas-go-modules/maincomposer/util/alog"
)

// Resolver is the resolver root.
type Resolver struct {
	BookClient controllers.BookController
}

// NewSchema creates a graphql executable schema.
func NewSchema(logger alog.Logger, client *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			BookClient: controllers.NewBookController(logger, client),
		},
	})
}
