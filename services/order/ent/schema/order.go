package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

const (
	_APPROVAL_PENDING string = "APPROVAL_PENDING"
	_APPROVED         string = "APPROVED"
)

// Annotations of the Order.
func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "orders"},
	}
}

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("request_id", uuid.UUID{}),
		field.String("user_id").NotEmpty(),
		field.String("address").NotEmpty(),
		field.Enum("status").Values(_APPROVAL_PENDING, _APPROVED),
		field.Time("update_at").
			Default(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
