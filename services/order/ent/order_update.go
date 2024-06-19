// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ntc-goer/microservice-examples/orderservice/ent/order"
	"github.com/ntc-goer/microservice-examples/orderservice/ent/predicate"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetRequestID sets the "request_id" field.
func (ou *OrderUpdate) SetRequestID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetRequestID(u)
	return ou
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableRequestID(u *uuid.UUID) *OrderUpdate {
	if u != nil {
		ou.SetRequestID(*u)
	}
	return ou
}

// SetUserID sets the "user_id" field.
func (ou *OrderUpdate) SetUserID(s string) *OrderUpdate {
	ou.mutation.SetUserID(s)
	return ou
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableUserID(s *string) *OrderUpdate {
	if s != nil {
		ou.SetUserID(*s)
	}
	return ou
}

// SetAddress sets the "address" field.
func (ou *OrderUpdate) SetAddress(s string) *OrderUpdate {
	ou.mutation.SetAddress(s)
	return ou
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableAddress(s *string) *OrderUpdate {
	if s != nil {
		ou.SetAddress(*s)
	}
	return ou
}

// SetStatus sets the "status" field.
func (ou *OrderUpdate) SetStatus(o order.Status) *OrderUpdate {
	ou.mutation.SetStatus(o)
	return ou
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableStatus(o *order.Status) *OrderUpdate {
	if o != nil {
		ou.SetStatus(*o)
	}
	return ou
}

// SetUpdateAt sets the "update_at" field.
func (ou *OrderUpdate) SetUpdateAt(t time.Time) *OrderUpdate {
	ou.mutation.SetUpdateAt(t)
	return ou
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableUpdateAt(t *time.Time) *OrderUpdate {
	if t != nil {
		ou.SetUpdateAt(*t)
	}
	return ou
}

// SetCreatedAt sets the "created_at" field.
func (ou *OrderUpdate) SetCreatedAt(t time.Time) *OrderUpdate {
	ou.mutation.SetCreatedAt(t)
	return ou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableCreatedAt(t *time.Time) *OrderUpdate {
	if t != nil {
		ou.SetCreatedAt(*t)
	}
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.UserID(); ok {
		if err := order.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Order.user_id": %w`, err)}
		}
	}
	if v, ok := ou.mutation.Address(); ok {
		if err := order.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Order.address": %w`, err)}
		}
	}
	if v, ok := ou.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.RequestID(); ok {
		_spec.SetField(order.FieldRequestID, field.TypeUUID, value)
	}
	if value, ok := ou.mutation.UserID(); ok {
		_spec.SetField(order.FieldUserID, field.TypeString, value)
	}
	if value, ok := ou.mutation.Address(); ok {
		_spec.SetField(order.FieldAddress, field.TypeString, value)
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ou.mutation.UpdateAt(); ok {
		_spec.SetField(order.FieldUpdateAt, field.TypeTime, value)
	}
	if value, ok := ou.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetRequestID sets the "request_id" field.
func (ouo *OrderUpdateOne) SetRequestID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetRequestID(u)
	return ouo
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableRequestID(u *uuid.UUID) *OrderUpdateOne {
	if u != nil {
		ouo.SetRequestID(*u)
	}
	return ouo
}

// SetUserID sets the "user_id" field.
func (ouo *OrderUpdateOne) SetUserID(s string) *OrderUpdateOne {
	ouo.mutation.SetUserID(s)
	return ouo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableUserID(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetUserID(*s)
	}
	return ouo
}

// SetAddress sets the "address" field.
func (ouo *OrderUpdateOne) SetAddress(s string) *OrderUpdateOne {
	ouo.mutation.SetAddress(s)
	return ouo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableAddress(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetAddress(*s)
	}
	return ouo
}

// SetStatus sets the "status" field.
func (ouo *OrderUpdateOne) SetStatus(o order.Status) *OrderUpdateOne {
	ouo.mutation.SetStatus(o)
	return ouo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableStatus(o *order.Status) *OrderUpdateOne {
	if o != nil {
		ouo.SetStatus(*o)
	}
	return ouo
}

// SetUpdateAt sets the "update_at" field.
func (ouo *OrderUpdateOne) SetUpdateAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetUpdateAt(t)
	return ouo
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableUpdateAt(t *time.Time) *OrderUpdateOne {
	if t != nil {
		ouo.SetUpdateAt(*t)
	}
	return ouo
}

// SetCreatedAt sets the "created_at" field.
func (ouo *OrderUpdateOne) SetCreatedAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetCreatedAt(t)
	return ouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableCreatedAt(t *time.Time) *OrderUpdateOne {
	if t != nil {
		ouo.SetCreatedAt(*t)
	}
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ouo *OrderUpdateOne) Where(ps ...predicate.Order) *OrderUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.UserID(); ok {
		if err := order.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Order.user_id": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.Address(); ok {
		if err := order.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Order.address": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.RequestID(); ok {
		_spec.SetField(order.FieldRequestID, field.TypeUUID, value)
	}
	if value, ok := ouo.mutation.UserID(); ok {
		_spec.SetField(order.FieldUserID, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Address(); ok {
		_spec.SetField(order.FieldAddress, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ouo.mutation.UpdateAt(); ok {
		_spec.SetField(order.FieldUpdateAt, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
