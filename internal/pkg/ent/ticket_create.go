// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/users/internal/pkg/ent/ticket"
	"github.com/pepeunlimited/users/internal/pkg/ent/user"
)

// TicketCreate is the builder for creating a Ticket entity.
type TicketCreate struct {
	config
	token      *string
	created_at *time.Time
	expires_at *time.Time
	users      map[int]struct{}
}

// SetToken sets the token field.
func (tc *TicketCreate) SetToken(s string) *TicketCreate {
	tc.token = &s
	return tc
}

// SetCreatedAt sets the created_at field.
func (tc *TicketCreate) SetCreatedAt(t time.Time) *TicketCreate {
	tc.created_at = &t
	return tc
}

// SetExpiresAt sets the expires_at field.
func (tc *TicketCreate) SetExpiresAt(t time.Time) *TicketCreate {
	tc.expires_at = &t
	return tc
}

// SetUsersID sets the users edge to User by id.
func (tc *TicketCreate) SetUsersID(id int) *TicketCreate {
	if tc.users == nil {
		tc.users = make(map[int]struct{})
	}
	tc.users[id] = struct{}{}
	return tc
}

// SetNillableUsersID sets the users edge to User by id if the given value is not nil.
func (tc *TicketCreate) SetNillableUsersID(id *int) *TicketCreate {
	if id != nil {
		tc = tc.SetUsersID(*id)
	}
	return tc
}

// SetUsers sets the users edge to User.
func (tc *TicketCreate) SetUsers(u *User) *TicketCreate {
	return tc.SetUsersID(u.ID)
}

// Save creates the Ticket in the database.
func (tc *TicketCreate) Save(ctx context.Context) (*Ticket, error) {
	if tc.token == nil {
		return nil, errors.New("ent: missing required field \"token\"")
	}
	if err := ticket.TokenValidator(*tc.token); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"token\": %v", err)
	}
	if tc.created_at == nil {
		return nil, errors.New("ent: missing required field \"created_at\"")
	}
	if tc.expires_at == nil {
		return nil, errors.New("ent: missing required field \"expires_at\"")
	}
	if len(tc.users) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"users\"")
	}
	return tc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TicketCreate) SaveX(ctx context.Context) *Ticket {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TicketCreate) sqlSave(ctx context.Context) (*Ticket, error) {
	var (
		t     = &Ticket{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: ticket.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ticket.FieldID,
			},
		}
	)
	if value := tc.token; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: ticket.FieldToken,
		})
		t.Token = *value
	}
	if value := tc.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: ticket.FieldCreatedAt,
		})
		t.CreatedAt = *value
	}
	if value := tc.expires_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: ticket.FieldExpiresAt,
		})
		t.ExpiresAt = *value
	}
	if nodes := tc.users; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.UsersTable,
			Columns: []string{ticket.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}