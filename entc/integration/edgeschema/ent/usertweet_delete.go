// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/usertweet"
	"entgo.io/ent/schema/field"
)

// UserTweetDelete is the builder for deleting a UserTweet entity.
type UserTweetDelete struct {
	config
	hooks    []Hook
	mutation *UserTweetMutation
}

// Where appends a list predicates to the UserTweetDelete builder.
func (utd *UserTweetDelete) Where(ps ...predicate.UserTweet) *UserTweetDelete {
	utd.mutation.Where(ps...)
	return utd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (utd *UserTweetDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(utd.hooks) == 0 {
		affected, err = utd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserTweetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			utd.mutation = mutation
			affected, err = utd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(utd.hooks) - 1; i >= 0; i-- {
			if utd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = utd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, utd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (utd *UserTweetDelete) ExecX(ctx context.Context) int {
	n, err := utd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (utd *UserTweetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: usertweet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertweet.FieldID,
			},
		},
	}
	if ps := utd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, utd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// UserTweetDeleteOne is the builder for deleting a single UserTweet entity.
type UserTweetDeleteOne struct {
	utd *UserTweetDelete
}

// Exec executes the deletion query.
func (utdo *UserTweetDeleteOne) Exec(ctx context.Context) error {
	n, err := utdo.utd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usertweet.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (utdo *UserTweetDeleteOne) ExecX(ctx context.Context) {
	utdo.utd.ExecX(ctx)
}
