// Code generated by ent, DO NOT EDIT.

package ticket

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/ntc-goer/microservice-examples/kitchen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldID, id))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldOrderID, v))
}

// RequestID applies equality check predicate on the "request_id" field. It's identical to RequestIDEQ.
func RequestID(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldRequestID, v))
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldUpdateAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldCreatedAt, v))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...string) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...string) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldOrderID, vs...))
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldOrderID, v))
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldOrderID, v))
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldOrderID, v))
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldOrderID, v))
}

// OrderIDContains applies the Contains predicate on the "order_id" field.
func OrderIDContains(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldContains(FieldOrderID, v))
}

// OrderIDHasPrefix applies the HasPrefix predicate on the "order_id" field.
func OrderIDHasPrefix(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldHasPrefix(FieldOrderID, v))
}

// OrderIDHasSuffix applies the HasSuffix predicate on the "order_id" field.
func OrderIDHasSuffix(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldHasSuffix(FieldOrderID, v))
}

// OrderIDEqualFold applies the EqualFold predicate on the "order_id" field.
func OrderIDEqualFold(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEqualFold(FieldOrderID, v))
}

// OrderIDContainsFold applies the ContainsFold predicate on the "order_id" field.
func OrderIDContainsFold(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldContainsFold(FieldOrderID, v))
}

// RequestIDEQ applies the EQ predicate on the "request_id" field.
func RequestIDEQ(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldRequestID, v))
}

// RequestIDNEQ applies the NEQ predicate on the "request_id" field.
func RequestIDNEQ(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldRequestID, v))
}

// RequestIDIn applies the In predicate on the "request_id" field.
func RequestIDIn(vs ...string) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldRequestID, vs...))
}

// RequestIDNotIn applies the NotIn predicate on the "request_id" field.
func RequestIDNotIn(vs ...string) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldRequestID, vs...))
}

// RequestIDGT applies the GT predicate on the "request_id" field.
func RequestIDGT(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldRequestID, v))
}

// RequestIDGTE applies the GTE predicate on the "request_id" field.
func RequestIDGTE(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldRequestID, v))
}

// RequestIDLT applies the LT predicate on the "request_id" field.
func RequestIDLT(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldRequestID, v))
}

// RequestIDLTE applies the LTE predicate on the "request_id" field.
func RequestIDLTE(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldRequestID, v))
}

// RequestIDContains applies the Contains predicate on the "request_id" field.
func RequestIDContains(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldContains(FieldRequestID, v))
}

// RequestIDHasPrefix applies the HasPrefix predicate on the "request_id" field.
func RequestIDHasPrefix(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldHasPrefix(FieldRequestID, v))
}

// RequestIDHasSuffix applies the HasSuffix predicate on the "request_id" field.
func RequestIDHasSuffix(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldHasSuffix(FieldRequestID, v))
}

// RequestIDEqualFold applies the EqualFold predicate on the "request_id" field.
func RequestIDEqualFold(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldEqualFold(FieldRequestID, v))
}

// RequestIDContainsFold applies the ContainsFold predicate on the "request_id" field.
func RequestIDContainsFold(v string) predicate.Ticket {
	return predicate.Ticket(sql.FieldContainsFold(FieldRequestID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldStatus, vs...))
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldUpdateAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Ticket {
	return predicate.Ticket(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ticket) predicate.Ticket {
	return predicate.Ticket(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ticket) predicate.Ticket {
	return predicate.Ticket(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Ticket) predicate.Ticket {
	return predicate.Ticket(sql.NotPredicates(p))
}