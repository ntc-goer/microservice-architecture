package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Dish holds the schema definition for the Dish entity.
type Dish struct {
	ent.Schema
}

// Annotations of the Order.
func (Dish) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "dishes"},
	}
}

// Fields of the Order.
func (Dish) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("order_id", uuid.UUID{}),
		field.String("dish_id").NotEmpty(),
		field.String("dish_name").NotEmpty(),
		field.Int("quantity").Min(1),
		field.Time("update_at").
			Default(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Order.
func (Dish) Edges() []ent.Edge {
	return nil
}
