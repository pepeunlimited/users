package schema

import (
	"errors"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"regexp"
)

var (
	ErrNotValidEmail = errors.New("ent: not valid email address")
)

const (
	EmailRegExp string = `^[a-zA-Z0-9.!#$%&'*+/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
)


// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Config() ent.Config {
	return ent.Config{
		Table: "users",
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").MaxLen(320).NotEmpty().Unique(),
		field.String("email").MaxLen(320).NotEmpty().Unique().Validate(func(s string) error {
			regex := regexp.MustCompile(EmailRegExp)
			if !regex.MatchString(s) {
				return ErrNotValidEmail
			}
			return nil
		}),
		field.String("password").MaxLen(72).Sensitive().NotEmpty(),
		field.Bool("is_deleted").Default(false),
		field.Bool("is_banned").Default(false),
		field.Bool("is_locked").Default(false),
		field.Time("last_modified"),
		field.Int64("profile_picture_id").Optional().Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tickets", Ticket.Type), // one-to-many
		edge.To("roles", Role.Type), 		// one-to-many
	}
}