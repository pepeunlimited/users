// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Role holds the value of the "role" field.
	Role string `json:"role,omitempty"`
}

// FromRows scans the sql response data into Role.
func (r *Role) FromRows(rows *sql.Rows) error {
	var scanr struct {
		ID   int
		Role sql.NullString
	}
	// the order here should be the same as in the `role.Columns`.
	if err := rows.Scan(
		&scanr.ID,
		&scanr.Role,
	); err != nil {
		return err
	}
	r.ID = scanr.ID
	r.Role = scanr.Role.String
	return nil
}

// QueryUsers queries the users edge of the Role.
func (r *Role) QueryUsers() *UserQuery {
	return (&RoleClient{r.config}).QueryUsers(r)
}

// Update returns a builder for updating this Role.
// Note that, you need to call Role.Unwrap() before calling this method, if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return (&RoleClient{r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", role=")
	builder.WriteString(r.Role)
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role

// FromRows scans the sql response data into Roles.
func (r *Roles) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scanr := &Role{}
		if err := scanr.FromRows(rows); err != nil {
			return err
		}
		*r = append(*r, scanr)
	}
	return nil
}

func (r Roles) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
