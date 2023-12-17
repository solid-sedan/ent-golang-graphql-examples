package controllers

import (
	"book-catalog/ent"
	"context"

	"github.com/atlas-health/atlas-go-modules/maincomposer/util/alog"
)

type BookController interface {
	// Mutations
	CreateAuthor(ctx context.Context, input ent.CreateAuthorInput) (*ent.Author, error)
	UpdateAuthor(ctx context.Context, id int, input ent.UpdateAuthorInput) (*ent.Author, error)

	CreateBook(ctx context.Context, input ent.CreateBookInput) (*ent.Book, error)
	UpdateBook(ctx context.Context, id int, input ent.UpdateBookInput) (*ent.Book, error)

	// Query
	Authors(ctx context.Context) ([]*ent.Author, error)
	Books(ctx context.Context) ([]*ent.Book, error)
}

type bookController struct {
	logger alog.Logger
	ent    *ent.Client
}

func NewBookController(logger alog.Logger, ent *ent.Client) BookController {
	return &bookController{
		logger: logger.As("struct", "bookController"),
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

func (c *bookController) Authors(ctx context.Context) ([]*ent.Author, error) {
	c.logger.Debug().Msg("Authors")
	return c.ent.Author.Query().All(ctx)
}

func (c *bookController) Books(ctx context.Context) ([]*ent.Book, error) {
	c.logger.Debug().Msg("Books")
	return c.ent.Book.Query().All(ctx)
}
