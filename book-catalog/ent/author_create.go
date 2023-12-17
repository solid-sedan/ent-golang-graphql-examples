// Code generated by ent, DO NOT EDIT.

package ent

import (
	"book-catalog/ent/author"
	"book-catalog/ent/book"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthorCreate is the builder for creating a Author entity.
type AuthorCreate struct {
	config
	mutation *AuthorMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AuthorCreate) SetName(s string) *AuthorCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetEmail sets the "email" field.
func (ac *AuthorCreate) SetEmail(s string) *AuthorCreate {
	ac.mutation.SetEmail(s)
	return ac
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (ac *AuthorCreate) AddBookIDs(ids ...int) *AuthorCreate {
	ac.mutation.AddBookIDs(ids...)
	return ac
}

// AddBooks adds the "books" edges to the Book entity.
func (ac *AuthorCreate) AddBooks(b ...*Book) *AuthorCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ac.AddBookIDs(ids...)
}

// Mutation returns the AuthorMutation object of the builder.
func (ac *AuthorCreate) Mutation() *AuthorMutation {
	return ac.mutation
}

// Save creates the Author in the database.
func (ac *AuthorCreate) Save(ctx context.Context) (*Author, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AuthorCreate) SaveX(ctx context.Context) *Author {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AuthorCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AuthorCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AuthorCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Author.name"`)}
	}
	if _, ok := ac.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Author.email"`)}
	}
	return nil
}

func (ac *AuthorCreate) sqlSave(ctx context.Context) (*Author, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AuthorCreate) createSpec() (*Author, *sqlgraph.CreateSpec) {
	var (
		_node = &Author{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(author.Table, sqlgraph.NewFieldSpec(author.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(author.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Email(); ok {
		_spec.SetField(author.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if nodes := ac.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   author.BooksTable,
			Columns: []string{author.BooksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AuthorCreateBulk is the builder for creating many Author entities in bulk.
type AuthorCreateBulk struct {
	config
	err      error
	builders []*AuthorCreate
}

// Save creates the Author entities in the database.
func (acb *AuthorCreateBulk) Save(ctx context.Context) ([]*Author, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Author, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AuthorCreateBulk) SaveX(ctx context.Context) []*Author {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AuthorCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AuthorCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
