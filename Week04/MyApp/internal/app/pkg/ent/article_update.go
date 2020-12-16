// Code generated by entc, DO NOT EDIT.

package ent

import (
	"MyApp/internal/app/pkg/ent/article"
	"MyApp/internal/app/pkg/ent/predicate"
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ArticleUpdate is the builder for updating Article entities.
type ArticleUpdate struct {
	config
	hooks    []Hook
	mutation *ArticleMutation
}

// Where adds a new predicate for the builder.
func (au *ArticleUpdate) Where(ps ...predicate.Article) *ArticleUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetTitle sets the title field.
func (au *ArticleUpdate) SetTitle(s string) *ArticleUpdate {
	au.mutation.SetTitle(s)
	return au
}

// SetBody sets the body field.
func (au *ArticleUpdate) SetBody(s string) *ArticleUpdate {
	au.mutation.SetBody(s)
	return au
}

// SetPublished sets the published field.
func (au *ArticleUpdate) SetPublished(b bool) *ArticleUpdate {
	au.mutation.SetPublished(b)
	return au
}

// SetCreatedTime sets the created_time field.
func (au *ArticleUpdate) SetCreatedTime(t time.Time) *ArticleUpdate {
	au.mutation.SetCreatedTime(t)
	return au
}

// SetNillableCreatedTime sets the created_time field if the given value is not nil.
func (au *ArticleUpdate) SetNillableCreatedTime(t *time.Time) *ArticleUpdate {
	if t != nil {
		au.SetCreatedTime(*t)
	}
	return au
}

// Mutation returns the ArticleMutation object of the builder.
func (au *ArticleUpdate) Mutation() *ArticleMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArticleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArticleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArticleUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArticleUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArticleUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ArticleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   article.Table,
			Columns: article.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: article.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldTitle,
		})
	}
	if value, ok := au.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldBody,
		})
	}
	if value, ok := au.mutation.Published(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: article.FieldPublished,
		})
	}
	if value, ok := au.mutation.CreatedTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: article.FieldCreatedTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ArticleUpdateOne is the builder for updating a single Article entity.
type ArticleUpdateOne struct {
	config
	hooks    []Hook
	mutation *ArticleMutation
}

// SetTitle sets the title field.
func (auo *ArticleUpdateOne) SetTitle(s string) *ArticleUpdateOne {
	auo.mutation.SetTitle(s)
	return auo
}

// SetBody sets the body field.
func (auo *ArticleUpdateOne) SetBody(s string) *ArticleUpdateOne {
	auo.mutation.SetBody(s)
	return auo
}

// SetPublished sets the published field.
func (auo *ArticleUpdateOne) SetPublished(b bool) *ArticleUpdateOne {
	auo.mutation.SetPublished(b)
	return auo
}

// SetCreatedTime sets the created_time field.
func (auo *ArticleUpdateOne) SetCreatedTime(t time.Time) *ArticleUpdateOne {
	auo.mutation.SetCreatedTime(t)
	return auo
}

// SetNillableCreatedTime sets the created_time field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableCreatedTime(t *time.Time) *ArticleUpdateOne {
	if t != nil {
		auo.SetCreatedTime(*t)
	}
	return auo
}

// Mutation returns the ArticleMutation object of the builder.
func (auo *ArticleUpdateOne) Mutation() *ArticleMutation {
	return auo.mutation
}

// Save executes the query and returns the updated entity.
func (auo *ArticleUpdateOne) Save(ctx context.Context) (*Article, error) {
	var (
		err  error
		node *Article
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArticleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArticleUpdateOne) SaveX(ctx context.Context) *Article {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArticleUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArticleUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ArticleUpdateOne) sqlSave(ctx context.Context) (_node *Article, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   article.Table,
			Columns: article.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: article.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Article.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldTitle,
		})
	}
	if value, ok := auo.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldBody,
		})
	}
	if value, ok := auo.mutation.Published(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: article.FieldPublished,
		})
	}
	if value, ok := auo.mutation.CreatedTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: article.FieldCreatedTime,
		})
	}
	_node = &Article{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
