// Code generated by ent, DO NOT EDIT.

package ent

import (
	"book-catalog/ent/author"
	"book-catalog/ent/book"
	"context"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// AuthorEdge is the edge representation of Author.
type AuthorEdge struct {
	Node   *Author `json:"node"`
	Cursor Cursor  `json:"cursor"`
}

// AuthorConnection is the connection containing edges to Author.
type AuthorConnection struct {
	Edges      []*AuthorEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

func (c *AuthorConnection) build(nodes []*Author, pager *authorPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Author
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Author {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Author {
			return nodes[i]
		}
	}
	c.Edges = make([]*AuthorEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &AuthorEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// AuthorPaginateOption enables pagination customization.
type AuthorPaginateOption func(*authorPager) error

// WithAuthorOrder configures pagination ordering.
func WithAuthorOrder(order *AuthorOrder) AuthorPaginateOption {
	if order == nil {
		order = DefaultAuthorOrder
	}
	o := *order
	return func(pager *authorPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultAuthorOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithAuthorFilter configures pagination filter.
func WithAuthorFilter(filter func(*AuthorQuery) (*AuthorQuery, error)) AuthorPaginateOption {
	return func(pager *authorPager) error {
		if filter == nil {
			return errors.New("AuthorQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type authorPager struct {
	reverse bool
	order   *AuthorOrder
	filter  func(*AuthorQuery) (*AuthorQuery, error)
}

func newAuthorPager(opts []AuthorPaginateOption, reverse bool) (*authorPager, error) {
	pager := &authorPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultAuthorOrder
	}
	return pager, nil
}

func (p *authorPager) applyFilter(query *AuthorQuery) (*AuthorQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *authorPager) toCursor(a *Author) Cursor {
	return p.order.Field.toCursor(a)
}

func (p *authorPager) applyCursors(query *AuthorQuery, after, before *Cursor) (*AuthorQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultAuthorOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *authorPager) applyOrder(query *AuthorQuery) *AuthorQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultAuthorOrder.Field {
		query = query.Order(DefaultAuthorOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *authorPager) orderExpr(query *AuthorQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultAuthorOrder.Field {
			b.Comma().Ident(DefaultAuthorOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Author.
func (a *AuthorQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...AuthorPaginateOption,
) (*AuthorConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newAuthorPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if a, err = pager.applyFilter(a); err != nil {
		return nil, err
	}
	conn := &AuthorConnection{Edges: []*AuthorEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = a.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if a, err = pager.applyCursors(a, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		a.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := a.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	a = pager.applyOrder(a)
	nodes, err := a.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// AuthorOrderField defines the ordering field of Author.
type AuthorOrderField struct {
	// Value extracts the ordering value from the given Author.
	Value    func(*Author) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) author.OrderOption
	toCursor func(*Author) Cursor
}

// AuthorOrder defines the ordering of Author.
type AuthorOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *AuthorOrderField `json:"field"`
}

// DefaultAuthorOrder is the default ordering of Author.
var DefaultAuthorOrder = &AuthorOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &AuthorOrderField{
		Value: func(a *Author) (ent.Value, error) {
			return a.ID, nil
		},
		column: author.FieldID,
		toTerm: author.ByID,
		toCursor: func(a *Author) Cursor {
			return Cursor{ID: a.ID}
		},
	},
}

// ToEdge converts Author into AuthorEdge.
func (a *Author) ToEdge(order *AuthorOrder) *AuthorEdge {
	if order == nil {
		order = DefaultAuthorOrder
	}
	return &AuthorEdge{
		Node:   a,
		Cursor: order.Field.toCursor(a),
	}
}

// BookEdge is the edge representation of Book.
type BookEdge struct {
	Node   *Book  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// BookConnection is the connection containing edges to Book.
type BookConnection struct {
	Edges      []*BookEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *BookConnection) build(nodes []*Book, pager *bookPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Book
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Book {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Book {
			return nodes[i]
		}
	}
	c.Edges = make([]*BookEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &BookEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// BookPaginateOption enables pagination customization.
type BookPaginateOption func(*bookPager) error

// WithBookOrder configures pagination ordering.
func WithBookOrder(order *BookOrder) BookPaginateOption {
	if order == nil {
		order = DefaultBookOrder
	}
	o := *order
	return func(pager *bookPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultBookOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithBookFilter configures pagination filter.
func WithBookFilter(filter func(*BookQuery) (*BookQuery, error)) BookPaginateOption {
	return func(pager *bookPager) error {
		if filter == nil {
			return errors.New("BookQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type bookPager struct {
	reverse bool
	order   *BookOrder
	filter  func(*BookQuery) (*BookQuery, error)
}

func newBookPager(opts []BookPaginateOption, reverse bool) (*bookPager, error) {
	pager := &bookPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultBookOrder
	}
	return pager, nil
}

func (p *bookPager) applyFilter(query *BookQuery) (*BookQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *bookPager) toCursor(b *Book) Cursor {
	return p.order.Field.toCursor(b)
}

func (p *bookPager) applyCursors(query *BookQuery, after, before *Cursor) (*BookQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultBookOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *bookPager) applyOrder(query *BookQuery) *BookQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultBookOrder.Field {
		query = query.Order(DefaultBookOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *bookPager) orderExpr(query *BookQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultBookOrder.Field {
			b.Comma().Ident(DefaultBookOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Book.
func (b *BookQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...BookPaginateOption,
) (*BookConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newBookPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if b, err = pager.applyFilter(b); err != nil {
		return nil, err
	}
	conn := &BookConnection{Edges: []*BookEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = b.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if b, err = pager.applyCursors(b, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		b.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := b.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	b = pager.applyOrder(b)
	nodes, err := b.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// BookOrderField defines the ordering field of Book.
type BookOrderField struct {
	// Value extracts the ordering value from the given Book.
	Value    func(*Book) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) book.OrderOption
	toCursor func(*Book) Cursor
}

// BookOrder defines the ordering of Book.
type BookOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *BookOrderField `json:"field"`
}

// DefaultBookOrder is the default ordering of Book.
var DefaultBookOrder = &BookOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &BookOrderField{
		Value: func(b *Book) (ent.Value, error) {
			return b.ID, nil
		},
		column: book.FieldID,
		toTerm: book.ByID,
		toCursor: func(b *Book) Cursor {
			return Cursor{ID: b.ID}
		},
	},
}

// ToEdge converts Book into BookEdge.
func (b *Book) ToEdge(order *BookOrder) *BookEdge {
	if order == nil {
		order = DefaultBookOrder
	}
	return &BookEdge{
		Node:   b,
		Cursor: order.Field.toCursor(b),
	}
}
