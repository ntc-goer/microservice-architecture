package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// SagaLogs holds the schema definition for the SagaLogs entity.
type SagaLogs struct {
	ent.Schema
}

// Annotations of the Ticket.
func (SagaLogs) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "saga_logs"},
	}
}

// Fields of the SagaLogs.
func (SagaLogs) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("workflow_id").NotEmpty(),
		field.String("workflow_name").NotEmpty(),
		field.String("step_name").NotEmpty(),
		field.Time("update_at").
			Default(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}
