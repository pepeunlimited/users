// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
	user2 "github.com/pepeunlimited/users/internal/app/app1/ent/user"
)

var (
	// TicketsColumns holds the columns for the "tickets" table.
	TicketsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString, Unique: true, Size: 72},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "users_id", Type: field.TypeInt, Nullable: true},
	}
	// TicketsTable holds the schema information for the "tickets" table.
	TicketsTable = &schema.Table{
		Name:       "tickets",
		Columns:    TicketsColumns,
		PrimaryKey: []*schema.Column{TicketsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "tickets_users_tickets",
				Columns: []*schema.Column{TicketsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role", Type: field.TypeString, Default: user2.DefaultRole},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 320},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 320},
		{Name: "password", Type: field.TypeString, Size: 72},
		{Name: "is_deleted", Type: field.TypeBool, Default: user2.DefaultIsDeleted},
		{Name: "is_banned", Type: field.TypeBool, Default: user2.DefaultIsBanned},
		{Name: "is_locked", Type: field.TypeBool, Default: user2.DefaultIsLocked},
		{Name: "last_modified", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TicketsTable,
		UsersTable,
	}
)

func init() {
	TicketsTable.ForeignKeys[0].RefTable = UsersTable
}