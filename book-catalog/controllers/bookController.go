package controllers

import (
	"book-catalog/ent"
	"context"

	"github.com/rs/zerolog"
)

type BookController interface {
	// Mutations
	CreateAuthor(ctx context.Context, input ent.CreateAuthorInput) (*ent.Author, error)
	UpdateAuthor(ctx context.Context, id int, input ent.UpdateAuthorInput) (*ent.Author, error)

	CreateBook(ctx context.Context, input ent.CreateBookInput) (*ent.Book, error)
	UpdateBook(ctx context.Context, id int, input ent.UpdateBookInput) (*ent.Book, error)

	// Query
	Authors(ctx context.Context, where *ent.AuthorWhereInput) ([]*ent.Author, error)
	Books(ctx context.Context) ([]*ent.Book, error)
}

type bookController struct {
	logger zerolog.Logger
	ent    *ent.Client
}

func NewBookController(logger zerolog.Logger, ent *ent.Client) BookController {
	return &bookController{
		logger: logger,
		ent:    ent,
	}
}

func (c *bookController) CreateAuthor(ctx context.Context, input ent.CreateAuthorInput) (*ent.Author, error) {
	c.logger.Debug().Msg("CreateAuthor")
	return c.ent.Author.Create().SetInput(input).Save(ctx)
}

func (c *bookController) UpdateAuthor(ctx context.Context, id int, input ent.UpdateAuthorInput) (*ent.Author, error) {
	c.logger.Debug().Msg("UpdateAuthor")
	return c.ent.Author.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (c *bookController) CreateBook(ctx context.Context, input ent.CreateBookInput) (*ent.Book, error) {
	c.logger.Debug().Msg("CreateBook")
	return c.ent.Book.Create().SetInput(input).Save(ctx)
}

func (c *bookController) UpdateBook(ctx context.Context, id int, input ent.UpdateBookInput) (*ent.Book, error) {
	c.logger.Debug().Msg("UpdateBook")
	return c.ent.Book.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (c *bookController) Authors(ctx context.Context, where *ent.AuthorWhereInput) ([]*ent.Author, error) {
	c.logger.Debug().Msg("Authors")
	// example with filter
	// more examples (paginations, order, etc) https://entgo.io/docs/tutorial-todo-gql-filter-input/#custom-filters
	query := c.ent.Author.Query()
	query, err := where.Filter(query)
	if err != nil {
		return nil, err
	}
	return query.All(ctx)
}

func (c *bookController) Books(ctx context.Context) ([]*ent.Book, error) {
	c.logger.Debug().Msg("Books")
	return c.ent.Book.Query().All(ctx)
}
