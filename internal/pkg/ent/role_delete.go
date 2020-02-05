// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/users/internal/pkg/ent/predicate"
	"github.com/pepeunlimited/users/internal/pkg/ent/role"
)

// RoleDelete is the builder for deleting a Role entity.
type RoleDelete struct {
	config
	predicates []predicate.Role
}

// Where adds a new predicate to the delete builder.
func (rd *RoleDelete) Where(ps ...predicate.Role) *RoleDelete {
	rd.predicates = append(rd.predicates, ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *RoleDelete) Exec(ctx context.Context) (int, error) {
	return rd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *RoleDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *RoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: role.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: role.FieldID,
			},
		},
	}
	if ps := rd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
}

// RoleDeleteOne is the builder for deleting a single Role entity.
type RoleDeleteOne struct {
	rd *RoleDelete
}

// Exec executes the deletion query.
func (rdo *RoleDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{role.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *RoleDeleteOne) ExecX(ctx context.Context) {
	rdo.rd.ExecX(ctx)
}