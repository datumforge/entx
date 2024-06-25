// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/entx/vanilla/_example/ent/orgmembership"
	"github.com/datumforge/entx/vanilla/_example/ent/predicate"
)

// OrgMembershipDelete is the builder for deleting a OrgMembership entity.
type OrgMembershipDelete struct {
	config
	hooks    []Hook
	mutation *OrgMembershipMutation
}

// Where appends a list predicates to the OrgMembershipDelete builder.
func (omd *OrgMembershipDelete) Where(ps ...predicate.OrgMembership) *OrgMembershipDelete {
	omd.mutation.Where(ps...)
	return omd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (omd *OrgMembershipDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, omd.sqlExec, omd.mutation, omd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (omd *OrgMembershipDelete) ExecX(ctx context.Context) int {
	n, err := omd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (omd *OrgMembershipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orgmembership.Table, sqlgraph.NewFieldSpec(orgmembership.FieldID, field.TypeString))
	if ps := omd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, omd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	omd.mutation.done = true
	return affected, err
}

// OrgMembershipDeleteOne is the builder for deleting a single OrgMembership entity.
type OrgMembershipDeleteOne struct {
	omd *OrgMembershipDelete
}

// Where appends a list predicates to the OrgMembershipDelete builder.
func (omdo *OrgMembershipDeleteOne) Where(ps ...predicate.OrgMembership) *OrgMembershipDeleteOne {
	omdo.omd.mutation.Where(ps...)
	return omdo
}

// Exec executes the deletion query.
func (omdo *OrgMembershipDeleteOne) Exec(ctx context.Context) error {
	n, err := omdo.omd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orgmembership.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (omdo *OrgMembershipDeleteOne) ExecX(ctx context.Context) {
	if err := omdo.Exec(ctx); err != nil {
		panic(err)
	}
}
