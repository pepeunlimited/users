package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Unique().MaxLen(72).NotEmpty(),
		field.Time("created_at"),
		field.Time("expires_at"),
	}
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("tickets").Unique(), // many-to-one
	}
}
