// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	ticket2 "github.com/pepeunlimited/users/internal/app/app1/ent/ticket"
	user2 "github.com/pepeunlimited/users/internal/app/app1/ent/user"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	role          *string
	username      *string
	email         *string
	password      *string
	is_deleted    *bool
	is_banned     *bool
	is_locked     *bool
	last_modified *time.Time
	tickets       map[int]struct{}
}

// SetRole sets the role field.
func (uc *UserCreate) SetRole(s string) *UserCreate {
	uc.role = &s
	return uc
}

// SetNillableRole sets the role field if the given value is not nil.
func (uc *UserCreate) SetNillableRole(s *string) *UserCreate {
	if s != nil {
		uc.SetRole(*s)
	}
	return uc
}

// SetUsername sets the username field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.username = &s
	return uc
}

// SetEmail sets the email field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.email = &s
	return uc
}

// SetPassword sets the password field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.password = &s
	return uc
}

// SetIsDeleted sets the is_deleted field.
func (uc *UserCreate) SetIsDeleted(b bool) *UserCreate {
	uc.is_deleted = &b
	return uc
}

// SetNillableIsDeleted sets the is_deleted field if the given value is not nil.
func (uc *UserCreate) SetNillableIsDeleted(b *bool) *UserCreate {
	if b != nil {
		uc.SetIsDeleted(*b)
	}
	return uc
}

// SetIsBanned sets the is_banned field.
func (uc *UserCreate) SetIsBanned(b bool) *UserCreate {
	uc.is_banned = &b
	return uc
}

// SetNillableIsBanned sets the is_banned field if the given value is not nil.
func (uc *UserCreate) SetNillableIsBanned(b *bool) *UserCreate {
	if b != nil {
		uc.SetIsBanned(*b)
	}
	return uc
}

// SetIsLocked sets the is_locked field.
func (uc *UserCreate) SetIsLocked(b bool) *UserCreate {
	uc.is_locked = &b
	return uc
}

// SetNillableIsLocked sets the is_locked field if the given value is not nil.
func (uc *UserCreate) SetNillableIsLocked(b *bool) *UserCreate {
	if b != nil {
		uc.SetIsLocked(*b)
	}
	return uc
}

// SetLastModified sets the last_modified field.
func (uc *UserCreate) SetLastModified(t time.Time) *UserCreate {
	uc.last_modified = &t
	return uc
}

// AddTicketIDs adds the tickets edge to Ticket by ids.
func (uc *UserCreate) AddTicketIDs(ids ...int) *UserCreate {
	if uc.tickets == nil {
		uc.tickets = make(map[int]struct{})
	}
	for i := range ids {
		uc.tickets[ids[i]] = struct{}{}
	}
	return uc
}

// AddTickets adds the tickets edges to Ticket.
func (uc *UserCreate) AddTickets(t ...*Ticket) *UserCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddTicketIDs(ids...)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.role == nil {
		v := user2.DefaultRole
		uc.role = &v
	}
	if uc.username == nil {
		return nil, errors.New("ent: missing required field \"username\"")
	}
	if err := user2.UsernameValidator(*uc.username); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"username\": %v", err)
	}
	if uc.email == nil {
		return nil, errors.New("ent: missing required field \"email\"")
	}
	if err := user2.EmailValidator(*uc.email); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"email\": %v", err)
	}
	if uc.password == nil {
		return nil, errors.New("ent: missing required field \"password\"")
	}
	if err := user2.PasswordValidator(*uc.password); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"password\": %v", err)
	}
	if uc.is_deleted == nil {
		v := user2.DefaultIsDeleted
		uc.is_deleted = &v
	}
	if uc.is_banned == nil {
		v := user2.DefaultIsBanned
		uc.is_banned = &v
	}
	if uc.is_locked == nil {
		v := user2.DefaultIsLocked
		uc.is_locked = &v
	}
	if uc.last_modified == nil {
		return nil, errors.New("ent: missing required field \"last_modified\"")
	}
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(uc.driver.Dialect())
		u       = &User{config: uc.config}
	)
	tx, err := uc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(user2.Table).Default()
	if value := uc.role; value != nil {
		insert.Set(user2.FieldRole, *value)
		u.Role = *value
	}
	if value := uc.username; value != nil {
		insert.Set(user2.FieldUsername, *value)
		u.Username = *value
	}
	if value := uc.email; value != nil {
		insert.Set(user2.FieldEmail, *value)
		u.Email = *value
	}
	if value := uc.password; value != nil {
		insert.Set(user2.FieldPassword, *value)
		u.Password = *value
	}
	if value := uc.is_deleted; value != nil {
		insert.Set(user2.FieldIsDeleted, *value)
		u.IsDeleted = *value
	}
	if value := uc.is_banned; value != nil {
		insert.Set(user2.FieldIsBanned, *value)
		u.IsBanned = *value
	}
	if value := uc.is_locked; value != nil {
		insert.Set(user2.FieldIsLocked, *value)
		u.IsLocked = *value
	}
	if value := uc.last_modified; value != nil {
		insert.Set(user2.FieldLastModified, *value)
		u.LastModified = *value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(user2.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	u.ID = int(id)
	if len(uc.tickets) > 0 {
		p := sql.P()
		for eid := range uc.tickets {
			p.Or().EQ(ticket2.FieldID, eid)
		}
		query, args := builder.Update(user2.TicketsTable).
			Set(user2.TicketsColumn, id).
			Where(sql.And(p, sql.IsNull(user2.TicketsColumn))).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(uc.tickets) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"tickets\" %v already connected to a different \"User\"", keys(uc.tickets))})
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}