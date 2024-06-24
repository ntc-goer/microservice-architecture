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
	_CREATE_PENDING      string = "CREATE_PENDING"
	_AWAITING_ACCEPTANCE string = "AWAITING_ACCEPTANCE"
	_CANCELED            string = "CANCELED"
)

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Annotations of the Ticket.
func (Ticket) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tickets"},
	}
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("order_id").NotEmpty(),
		field.String("request_id").NotEmpty(),
		field.Enum("status").Values(_CREATE_PENDING, _AWAITING_ACCEPTANCE, _CANCELED),
		field.Time("update_at").
			Default(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}
