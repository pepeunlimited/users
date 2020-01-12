package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Role struct {
	ent.Schema
}


func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("role").MaxLen(200).NotEmpty().Default("user"),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("roles").Unique(), // many-to-one
	}
}