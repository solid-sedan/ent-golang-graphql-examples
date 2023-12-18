package controllers

import (
	"book-catalog/ent"
	"context"

	"entgo.io/contrib/entgql"
	"github.com/rs/zerolog"
)

type BookController interface {
	// Mutations
	CreateAuthor(ctx context.Context, input ent.CreateAuthorInput) (*ent.Author, error)
	UpdateAuthor(ctx context.Context, id int, input ent.UpdateAuthorInput) (*ent.Author, error)

	CreateBook(ctx context.Context, input ent.CreateBookInput) (*ent.Book, error)
	UpdateBook(ctx context.Context, id int, input ent.UpdateBookInput) (*ent.Book, error)

	// Queries
	Authors(
		ctx context.Context,
		after *entgql.Cursor[int],
		first *int,
		before *entgql.Cursor[int],
		last *int,
		where *ent.AuthorWhereInput,
	) (*ent.AuthorConnection, error)
	Books(
		ctx context.Context,
		after *entgql.Cursor[int],
		first *int,
		before *entgql.Cursor[int],
		last *int,
		where *ent.BookWhereInput,
	) (*ent.BookConnection, error)
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

func (c *bookController) Authors(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.AuthorWhereInput) (*ent.AuthorConnection, error) {
	c.logger.Debug().Msg("Authors")
	// more examples (paginations, order, etc) https://entgo.io/docs/tutorial-todo-gql-filter-input/#custom-filters
	return c.ent.Author.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithAuthorFilter(where.Filter),
		)
}

func (c *bookController) Books(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.BookWhereInput) (*ent.BookConnection, error) {
	c.logger.Debug().Msg("Books")
	return c.ent.Book.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithBookFilter(where.Filter),
		)
}
