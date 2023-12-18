// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"book-catalog/ent"
	"bytes"
	"context"
	"errors"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Author struct {
		Books func(childComplexity int) int
		Email func(childComplexity int) int
		ID    func(childComplexity int) int
		Name  func(childComplexity int) int
	}

	AuthorConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	AuthorEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	Book struct {
		Author          func(childComplexity int) int
		CreatedAt       func(childComplexity int) int
		Genre           func(childComplexity int) int
		ID              func(childComplexity int) int
		Isbn            func(childComplexity int) int
		PublicationDate func(childComplexity int) int
		Title           func(childComplexity int) int
	}

	BookConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	BookEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	Mutation struct {
		CreateAuthor func(childComplexity int, input ent.CreateAuthorInput) int
		CreateBook   func(childComplexity int, input ent.CreateBookInput) int
		UpdateAuthor func(childComplexity int, id int, input ent.UpdateAuthorInput) int
		UpdateBook   func(childComplexity int, id int, input ent.UpdateBookInput) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Query struct {
		Authors func(childComplexity int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.AuthorWhereInput) int
		Books   func(childComplexity int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.BookWhereInput) int
		Node    func(childComplexity int, id int) int
		Nodes   func(childComplexity int, ids []int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Author.books":
		if e.complexity.Author.Books == nil {
			break
		}

		return e.complexity.Author.Books(childComplexity), true

	case "Author.email":
		if e.complexity.Author.Email == nil {
			break
		}

		return e.complexity.Author.Email(childComplexity), true

	case "Author.id":
		if e.complexity.Author.ID == nil {
			break
		}

		return e.complexity.Author.ID(childComplexity), true

	case "Author.name":
		if e.complexity.Author.Name == nil {
			break
		}

		return e.complexity.Author.Name(childComplexity), true

	case "AuthorConnection.edges":
		if e.complexity.AuthorConnection.Edges == nil {
			break
		}

		return e.complexity.AuthorConnection.Edges(childComplexity), true

	case "AuthorConnection.pageInfo":
		if e.complexity.AuthorConnection.PageInfo == nil {
			break
		}

		return e.complexity.AuthorConnection.PageInfo(childComplexity), true

	case "AuthorConnection.totalCount":
		if e.complexity.AuthorConnection.TotalCount == nil {
			break
		}

		return e.complexity.AuthorConnection.TotalCount(childComplexity), true

	case "AuthorEdge.cursor":
		if e.complexity.AuthorEdge.Cursor == nil {
			break
		}

		return e.complexity.AuthorEdge.Cursor(childComplexity), true

	case "AuthorEdge.node":
		if e.complexity.AuthorEdge.Node == nil {
			break
		}

		return e.complexity.AuthorEdge.Node(childComplexity), true

	case "Book.author":
		if e.complexity.Book.Author == nil {
			break
		}

		return e.complexity.Book.Author(childComplexity), true

	case "Book.createdAt":
		if e.complexity.Book.CreatedAt == nil {
			break
		}

		return e.complexity.Book.CreatedAt(childComplexity), true

	case "Book.genre":
		if e.complexity.Book.Genre == nil {
			break
		}

		return e.complexity.Book.Genre(childComplexity), true

	case "Book.id":
		if e.complexity.Book.ID == nil {
			break
		}

		return e.complexity.Book.ID(childComplexity), true

	case "Book.isbn":
		if e.complexity.Book.Isbn == nil {
			break
		}

		return e.complexity.Book.Isbn(childComplexity), true

	case "Book.publicationDate":
		if e.complexity.Book.PublicationDate == nil {
			break
		}

		return e.complexity.Book.PublicationDate(childComplexity), true

	case "Book.title":
		if e.complexity.Book.Title == nil {
			break
		}

		return e.complexity.Book.Title(childComplexity), true

	case "BookConnection.edges":
		if e.complexity.BookConnection.Edges == nil {
			break
		}

		return e.complexity.BookConnection.Edges(childComplexity), true

	case "BookConnection.pageInfo":
		if e.complexity.BookConnection.PageInfo == nil {
			break
		}

		return e.complexity.BookConnection.PageInfo(childComplexity), true

	case "BookConnection.totalCount":
		if e.complexity.BookConnection.TotalCount == nil {
			break
		}

		return e.complexity.BookConnection.TotalCount(childComplexity), true

	case "BookEdge.cursor":
		if e.complexity.BookEdge.Cursor == nil {
			break
		}

		return e.complexity.BookEdge.Cursor(childComplexity), true

	case "BookEdge.node":
		if e.complexity.BookEdge.Node == nil {
			break
		}

		return e.complexity.BookEdge.Node(childComplexity), true

	case "Mutation.createAuthor":
		if e.complexity.Mutation.CreateAuthor == nil {
			break
		}

		args, err := ec.field_Mutation_createAuthor_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateAuthor(childComplexity, args["input"].(ent.CreateAuthorInput)), true

	case "Mutation.createBook":
		if e.complexity.Mutation.CreateBook == nil {
			break
		}

		args, err := ec.field_Mutation_createBook_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateBook(childComplexity, args["input"].(ent.CreateBookInput)), true

	case "Mutation.updateAuthor":
		if e.complexity.Mutation.UpdateAuthor == nil {
			break
		}

		args, err := ec.field_Mutation_updateAuthor_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateAuthor(childComplexity, args["id"].(int), args["input"].(ent.UpdateAuthorInput)), true

	case "Mutation.updateBook":
		if e.complexity.Mutation.UpdateBook == nil {
			break
		}

		args, err := ec.field_Mutation_updateBook_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateBook(childComplexity, args["id"].(int), args["input"].(ent.UpdateBookInput)), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Query.authors":
		if e.complexity.Query.Authors == nil {
			break
		}

		args, err := ec.field_Query_authors_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Authors(childComplexity, args["after"].(*entgql.Cursor[int]), args["first"].(*int), args["before"].(*entgql.Cursor[int]), args["last"].(*int), args["where"].(*ent.AuthorWhereInput)), true

	case "Query.books":
		if e.complexity.Query.Books == nil {
			break
		}

		args, err := ec.field_Query_books_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Books(childComplexity, args["after"].(*entgql.Cursor[int]), args["first"].(*int), args["before"].(*entgql.Cursor[int]), args["last"].(*int), args["where"].(*ent.BookWhereInput)), true

	case "Query.node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(int)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]int)), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputAuthorWhereInput,
		ec.unmarshalInputBookWhereInput,
		ec.unmarshalInputCreateAuthorInput,
		ec.unmarshalInputCreateBookInput,
		ec.unmarshalInputUpdateAuthorInput,
		ec.unmarshalInputUpdateBookInput,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, rc.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../book.graphql", Input: `type Mutation {
    # The input and the output are types generated by Ent.
    createAuthor(input: CreateAuthorInput!): Author!
    updateAuthor(id: ID!, input: UpdateAuthorInput!): Author!
    
    createBook(input: CreateBookInput!): Book!
    updateBook(id: ID!, input: UpdateBookInput!): Book!
}`, BuiltIn: false},
	{Name: "../ent.graphql", Input: `directive @goField(forceResolver: Boolean, name: String) on FIELD_DEFINITION | INPUT_FIELD_DEFINITION
directive @goModel(model: String, models: [String!]) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION
type Author implements Node {
  id: ID!
  name: String!
  email: String!
  books: [Book!]
}
"""A connection to a list of items."""
type AuthorConnection {
  """A list of edges."""
  edges: [AuthorEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  """Identifies the total count of items in the connection."""
  totalCount: Int!
}
"""An edge in a connection."""
type AuthorEdge {
  """The item at the end of the edge."""
  node: Author
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""
AuthorWhereInput is used for filtering Author objects.
Input was generated by ent.
"""
input AuthorWhereInput {
  not: AuthorWhereInput
  and: [AuthorWhereInput!]
  or: [AuthorWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """email field predicates"""
  email: String
  emailNEQ: String
  emailIn: [String!]
  emailNotIn: [String!]
  emailGT: String
  emailGTE: String
  emailLT: String
  emailLTE: String
  emailContains: String
  emailHasPrefix: String
  emailHasSuffix: String
  emailEqualFold: String
  emailContainsFold: String
  """books edge predicates"""
  hasBooks: Boolean
  hasBooksWith: [BookWhereInput!]
}
type Book implements Node {
  id: ID!
  title: String!
  genre: String!
  publicationDate: String!
  isbn: String!
  createdAt: Time!
  author: Author
}
"""A connection to a list of items."""
type BookConnection {
  """A list of edges."""
  edges: [BookEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  """Identifies the total count of items in the connection."""
  totalCount: Int!
}
"""An edge in a connection."""
type BookEdge {
  """The item at the end of the edge."""
  node: Book
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""
BookWhereInput is used for filtering Book objects.
Input was generated by ent.
"""
input BookWhereInput {
  not: BookWhereInput
  and: [BookWhereInput!]
  or: [BookWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """title field predicates"""
  title: String
  titleNEQ: String
  titleIn: [String!]
  titleNotIn: [String!]
  titleGT: String
  titleGTE: String
  titleLT: String
  titleLTE: String
  titleContains: String
  titleHasPrefix: String
  titleHasSuffix: String
  titleEqualFold: String
  titleContainsFold: String
  """genre field predicates"""
  genre: String
  genreNEQ: String
  genreIn: [String!]
  genreNotIn: [String!]
  genreGT: String
  genreGTE: String
  genreLT: String
  genreLTE: String
  genreContains: String
  genreHasPrefix: String
  genreHasSuffix: String
  genreEqualFold: String
  genreContainsFold: String
  """publication_date field predicates"""
  publicationDate: String
  publicationDateNEQ: String
  publicationDateIn: [String!]
  publicationDateNotIn: [String!]
  publicationDateGT: String
  publicationDateGTE: String
  publicationDateLT: String
  publicationDateLTE: String
  publicationDateContains: String
  publicationDateHasPrefix: String
  publicationDateHasSuffix: String
  publicationDateEqualFold: String
  publicationDateContainsFold: String
  """isbn field predicates"""
  isbn: String
  isbnNEQ: String
  isbnIn: [String!]
  isbnNotIn: [String!]
  isbnGT: String
  isbnGTE: String
  isbnLT: String
  isbnLTE: String
  isbnContains: String
  isbnHasPrefix: String
  isbnHasSuffix: String
  isbnEqualFold: String
  isbnContainsFold: String
  """created_at field predicates"""
  createdAt: Time
  createdAtNEQ: Time
  createdAtIn: [Time!]
  createdAtNotIn: [Time!]
  createdAtGT: Time
  createdAtGTE: Time
  createdAtLT: Time
  createdAtLTE: Time
  """author edge predicates"""
  hasAuthor: Boolean
  hasAuthorWith: [AuthorWhereInput!]
}
"""
CreateAuthorInput is used for create Author object.
Input was generated by ent.
"""
input CreateAuthorInput {
  name: String!
  email: String!
  bookIDs: [ID!]
}
"""
CreateBookInput is used for create Book object.
Input was generated by ent.
"""
input CreateBookInput {
  title: String!
  genre: String!
  publicationDate: String!
  isbn: String!
  createdAt: Time
  authorID: ID
}
"""
Define a Relay Cursor type:
https://relay.dev/graphql/connections.htm#sec-Cursor
"""
scalar Cursor
"""
An object with an ID.
Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
"""
interface Node @goModel(model: "book-catalog/ent.Noder") {
  """The id of the object."""
  id: ID!
}
"""Possible directions in which to order a list of items when provided an ` + "`" + `orderBy` + "`" + ` argument."""
enum OrderDirection {
  """Specifies an ascending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  ASC
  """Specifies a descending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  DESC
}
"""
Information about pagination in a connection.
https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
"""
type PageInfo {
  """When paginating forwards, are there more items?"""
  hasNextPage: Boolean!
  """When paginating backwards, are there more items?"""
  hasPreviousPage: Boolean!
  """When paginating backwards, the cursor to continue."""
  startCursor: Cursor
  """When paginating forwards, the cursor to continue."""
  endCursor: Cursor
}
type Query {
  """Fetches an object given its ID."""
  node(
    """ID of the object."""
    id: ID!
  ): Node
  """Lookup nodes by a list of IDs."""
  nodes(
    """The list of node IDs."""
    ids: [ID!]!
  ): [Node]!
  authors(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """Filtering options for Authors returned from the connection."""
    where: AuthorWhereInput
  ): AuthorConnection!
  books(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """Filtering options for Books returned from the connection."""
    where: BookWhereInput
  ): BookConnection!
}
"""The builtin Time type"""
scalar Time
"""
UpdateAuthorInput is used for update Author object.
Input was generated by ent.
"""
input UpdateAuthorInput {
  name: String
  email: String
  addBookIDs: [ID!]
  removeBookIDs: [ID!]
  clearBooks: Boolean
}
"""
UpdateBookInput is used for update Book object.
Input was generated by ent.
"""
input UpdateBookInput {
  title: String
  genre: String
  publicationDate: String
  isbn: String
  createdAt: Time
  authorID: ID
  clearAuthor: Boolean
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
