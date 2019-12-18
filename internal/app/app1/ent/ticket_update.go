// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pepeunlimited/users/internal/app/app1/ent/predicate"
	"github.com/pepeunlimited/users/internal/app/app1/ent/ticket"
	"github.com/pepeunlimited/users/internal/app/app1/ent/user"
)

// TicketUpdate is the builder for updating Ticket entities.
type TicketUpdate struct {
	config
	token        *string
	created_at   *time.Time
	expires_at   *time.Time
	users        map[int]struct{}
	clearedUsers bool
	predicates   []predicate.Ticket
}

// Where adds a new predicate for the builder.
func (tu *TicketUpdate) Where(ps ...predicate.Ticket) *TicketUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetToken sets the token field.
func (tu *TicketUpdate) SetToken(s string) *TicketUpdate {
	tu.token = &s
	return tu
}

// SetCreatedAt sets the created_at field.
func (tu *TicketUpdate) SetCreatedAt(t time.Time) *TicketUpdate {
	tu.created_at = &t
	return tu
}

// SetExpiresAt sets the expires_at field.
func (tu *TicketUpdate) SetExpiresAt(t time.Time) *TicketUpdate {
	tu.expires_at = &t
	return tu
}

// SetUsersID sets the users edge to User by id.
func (tu *TicketUpdate) SetUsersID(id int) *TicketUpdate {
	if tu.users == nil {
		tu.users = make(map[int]struct{})
	}
	tu.users[id] = struct{}{}
	return tu
}

// SetNillableUsersID sets the users edge to User by id if the given value is not nil.
func (tu *TicketUpdate) SetNillableUsersID(id *int) *TicketUpdate {
	if id != nil {
		tu = tu.SetUsersID(*id)
	}
	return tu
}

// SetUsers sets the users edge to User.
func (tu *TicketUpdate) SetUsers(u *User) *TicketUpdate {
	return tu.SetUsersID(u.ID)
}

// ClearUsers clears the users edge to User.
func (tu *TicketUpdate) ClearUsers() *TicketUpdate {
	tu.clearedUsers = true
	return tu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TicketUpdate) Save(ctx context.Context) (int, error) {
	if tu.token != nil {
		if err := ticket.TokenValidator(*tu.token); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"token\": %v", err)
		}
	}
	if len(tu.users) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"users\"")
	}
	return tu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TicketUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TicketUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TicketUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TicketUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(tu.driver.Dialect())
		selector = builder.Select(ticket.FieldID).From(builder.Table(ticket.Table))
	)
	for _, p := range tu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = tu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := tu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(ticket.Table)
	)
	updater = updater.Where(sql.InInts(ticket.FieldID, ids...))
	if value := tu.token; value != nil {
		updater.Set(ticket.FieldToken, *value)
	}
	if value := tu.created_at; value != nil {
		updater.Set(ticket.FieldCreatedAt, *value)
	}
	if value := tu.expires_at; value != nil {
		updater.Set(ticket.FieldExpiresAt, *value)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if tu.clearedUsers {
		query, args := builder.Update(ticket.UsersTable).
			SetNull(ticket.UsersColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(tu.users) > 0 {
		for eid := range tu.users {
			query, args := builder.Update(ticket.UsersTable).
				Set(ticket.UsersColumn, eid).
				Where(sql.InInts(ticket.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// TicketUpdateOne is the builder for updating a single Ticket entity.
type TicketUpdateOne struct {
	config
	id           int
	token        *string
	created_at   *time.Time
	expires_at   *time.Time
	users        map[int]struct{}
	clearedUsers bool
}

// SetToken sets the token field.
func (tuo *TicketUpdateOne) SetToken(s string) *TicketUpdateOne {
	tuo.token = &s
	return tuo
}

// SetCreatedAt sets the created_at field.
func (tuo *TicketUpdateOne) SetCreatedAt(t time.Time) *TicketUpdateOne {
	tuo.created_at = &t
	return tuo
}

// SetExpiresAt sets the expires_at field.
func (tuo *TicketUpdateOne) SetExpiresAt(t time.Time) *TicketUpdateOne {
	tuo.expires_at = &t
	return tuo
}

// SetUsersID sets the users edge to User by id.
func (tuo *TicketUpdateOne) SetUsersID(id int) *TicketUpdateOne {
	if tuo.users == nil {
		tuo.users = make(map[int]struct{})
	}
	tuo.users[id] = struct{}{}
	return tuo
}

// SetNillableUsersID sets the users edge to User by id if the given value is not nil.
func (tuo *TicketUpdateOne) SetNillableUsersID(id *int) *TicketUpdateOne {
	if id != nil {
		tuo = tuo.SetUsersID(*id)
	}
	return tuo
}

// SetUsers sets the users edge to User.
func (tuo *TicketUpdateOne) SetUsers(u *User) *TicketUpdateOne {
	return tuo.SetUsersID(u.ID)
}

// ClearUsers clears the users edge to User.
func (tuo *TicketUpdateOne) ClearUsers() *TicketUpdateOne {
	tuo.clearedUsers = true
	return tuo
}

// Save executes the query and returns the updated entity.
func (tuo *TicketUpdateOne) Save(ctx context.Context) (*Ticket, error) {
	if tuo.token != nil {
		if err := ticket.TokenValidator(*tuo.token); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"token\": %v", err)
		}
	}
	if len(tuo.users) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"users\"")
	}
	return tuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TicketUpdateOne) SaveX(ctx context.Context) *Ticket {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TicketUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TicketUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TicketUpdateOne) sqlSave(ctx context.Context) (t *Ticket, err error) {
	var (
		builder  = sql.Dialect(tuo.driver.Dialect())
		selector = builder.Select(ticket.Columns...).From(builder.Table(ticket.Table))
	)
	ticket.ID(tuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = tuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		t = &Ticket{config: tuo.config}
		if err := t.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Ticket: %v", err)
		}
		id = t.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Ticket with id: %v", tuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Ticket with the same id: %v", tuo.id)
	}

	tx, err := tuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(ticket.Table)
	)
	updater = updater.Where(sql.InInts(ticket.FieldID, ids...))
	if value := tuo.token; value != nil {
		updater.Set(ticket.FieldToken, *value)
		t.Token = *value
	}
	if value := tuo.created_at; value != nil {
		updater.Set(ticket.FieldCreatedAt, *value)
		t.CreatedAt = *value
	}
	if value := tuo.expires_at; value != nil {
		updater.Set(ticket.FieldExpiresAt, *value)
		t.ExpiresAt = *value
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if tuo.clearedUsers {
		query, args := builder.Update(ticket.UsersTable).
			SetNull(ticket.UsersColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(tuo.users) > 0 {
		for eid := range tuo.users {
			query, args := builder.Update(ticket.UsersTable).
				Set(ticket.UsersColumn, eid).
				Where(sql.InInts(ticket.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return t, nil
}
