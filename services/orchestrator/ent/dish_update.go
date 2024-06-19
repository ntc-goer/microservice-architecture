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
	"github.com/ntc-goer/microservice-examples/orderservice/ent/dish"
	"github.com/ntc-goer/microservice-examples/orderservice/ent/predicate"
)

// DishUpdate is the builder for updating Dish entities.
type DishUpdate struct {
	config
	hooks    []Hook
	mutation *DishMutation
}

// Where appends a list predicates to the DishUpdate builder.
func (du *DishUpdate) Where(ps ...predicate.Dish) *DishUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetOrderID sets the "order_id" field.
func (du *DishUpdate) SetOrderID(u uuid.UUID) *DishUpdate {
	du.mutation.SetOrderID(u)
	return du
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (du *DishUpdate) SetNillableOrderID(u *uuid.UUID) *DishUpdate {
	if u != nil {
		du.SetOrderID(*u)
	}
	return du
}

// SetDishName sets the "dish_name" field.
func (du *DishUpdate) SetDishName(s string) *DishUpdate {
	du.mutation.SetDishName(s)
	return du
}

// SetNillableDishName sets the "dish_name" field if the given value is not nil.
func (du *DishUpdate) SetNillableDishName(s *string) *DishUpdate {
	if s != nil {
		du.SetDishName(*s)
	}
	return du
}

// SetQuantity sets the "quantity" field.
func (du *DishUpdate) SetQuantity(i int) *DishUpdate {
	du.mutation.ResetQuantity()
	du.mutation.SetQuantity(i)
	return du
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (du *DishUpdate) SetNillableQuantity(i *int) *DishUpdate {
	if i != nil {
		du.SetQuantity(*i)
	}
	return du
}

// AddQuantity adds i to the "quantity" field.
func (du *DishUpdate) AddQuantity(i int) *DishUpdate {
	du.mutation.AddQuantity(i)
	return du
}

// SetUpdateAt sets the "update_at" field.
func (du *DishUpdate) SetUpdateAt(t time.Time) *DishUpdate {
	du.mutation.SetUpdateAt(t)
	return du
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (du *DishUpdate) SetNillableUpdateAt(t *time.Time) *DishUpdate {
	if t != nil {
		du.SetUpdateAt(*t)
	}
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DishUpdate) SetCreatedAt(t time.Time) *DishUpdate {
	du.mutation.SetCreatedAt(t)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DishUpdate) SetNillableCreatedAt(t *time.Time) *DishUpdate {
	if t != nil {
		du.SetCreatedAt(*t)
	}
	return du
}

// Mutation returns the DishMutation object of the builder.
func (du *DishUpdate) Mutation() *DishMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DishUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DishUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DishUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DishUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DishUpdate) check() error {
	if v, ok := du.mutation.DishName(); ok {
		if err := dish.DishNameValidator(v); err != nil {
			return &ValidationError{Name: "dish_name", err: fmt.Errorf(`ent: validator failed for field "Dish.dish_name": %w`, err)}
		}
	}
	if v, ok := du.mutation.Quantity(); ok {
		if err := dish.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "Dish.quantity": %w`, err)}
		}
	}
	return nil
}

func (du *DishUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(dish.Table, dish.Columns, sqlgraph.NewFieldSpec(dish.FieldID, field.TypeUUID))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.OrderID(); ok {
		_spec.SetField(dish.FieldOrderID, field.TypeUUID, value)
	}
	if value, ok := du.mutation.DishName(); ok {
		_spec.SetField(dish.FieldDishName, field.TypeString, value)
	}
	if value, ok := du.mutation.Quantity(); ok {
		_spec.SetField(dish.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := du.mutation.AddedQuantity(); ok {
		_spec.AddField(dish.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := du.mutation.UpdateAt(); ok {
		_spec.SetField(dish.FieldUpdateAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.SetField(dish.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dish.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DishUpdateOne is the builder for updating a single Dish entity.
type DishUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DishMutation
}

// SetOrderID sets the "order_id" field.
func (duo *DishUpdateOne) SetOrderID(u uuid.UUID) *DishUpdateOne {
	duo.mutation.SetOrderID(u)
	return duo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (duo *DishUpdateOne) SetNillableOrderID(u *uuid.UUID) *DishUpdateOne {
	if u != nil {
		duo.SetOrderID(*u)
	}
	return duo
}

// SetDishName sets the "dish_name" field.
func (duo *DishUpdateOne) SetDishName(s string) *DishUpdateOne {
	duo.mutation.SetDishName(s)
	return duo
}

// SetNillableDishName sets the "dish_name" field if the given value is not nil.
func (duo *DishUpdateOne) SetNillableDishName(s *string) *DishUpdateOne {
	if s != nil {
		duo.SetDishName(*s)
	}
	return duo
}

// SetQuantity sets the "quantity" field.
func (duo *DishUpdateOne) SetQuantity(i int) *DishUpdateOne {
	duo.mutation.ResetQuantity()
	duo.mutation.SetQuantity(i)
	return duo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (duo *DishUpdateOne) SetNillableQuantity(i *int) *DishUpdateOne {
	if i != nil {
		duo.SetQuantity(*i)
	}
	return duo
}

// AddQuantity adds i to the "quantity" field.
func (duo *DishUpdateOne) AddQuantity(i int) *DishUpdateOne {
	duo.mutation.AddQuantity(i)
	return duo
}

// SetUpdateAt sets the "update_at" field.
func (duo *DishUpdateOne) SetUpdateAt(t time.Time) *DishUpdateOne {
	duo.mutation.SetUpdateAt(t)
	return duo
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (duo *DishUpdateOne) SetNillableUpdateAt(t *time.Time) *DishUpdateOne {
	if t != nil {
		duo.SetUpdateAt(*t)
	}
	return duo
}

// SetCreatedAt sets the "created_at" field.
func (duo *DishUpdateOne) SetCreatedAt(t time.Time) *DishUpdateOne {
	duo.mutation.SetCreatedAt(t)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DishUpdateOne) SetNillableCreatedAt(t *time.Time) *DishUpdateOne {
	if t != nil {
		duo.SetCreatedAt(*t)
	}
	return duo
}

// Mutation returns the DishMutation object of the builder.
func (duo *DishUpdateOne) Mutation() *DishMutation {
	return duo.mutation
}

// Where appends a list predicates to the DishUpdate builder.
func (duo *DishUpdateOne) Where(ps ...predicate.Dish) *DishUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DishUpdateOne) Select(field string, fields ...string) *DishUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Dish entity.
func (duo *DishUpdateOne) Save(ctx context.Context) (*Dish, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DishUpdateOne) SaveX(ctx context.Context) *Dish {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DishUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DishUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DishUpdateOne) check() error {
	if v, ok := duo.mutation.DishName(); ok {
		if err := dish.DishNameValidator(v); err != nil {
			return &ValidationError{Name: "dish_name", err: fmt.Errorf(`ent: validator failed for field "Dish.dish_name": %w`, err)}
		}
	}
	if v, ok := duo.mutation.Quantity(); ok {
		if err := dish.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "Dish.quantity": %w`, err)}
		}
	}
	return nil
}

func (duo *DishUpdateOne) sqlSave(ctx context.Context) (_node *Dish, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(dish.Table, dish.Columns, sqlgraph.NewFieldSpec(dish.FieldID, field.TypeUUID))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Dish.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dish.FieldID)
		for _, f := range fields {
			if !dish.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dish.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.OrderID(); ok {
		_spec.SetField(dish.FieldOrderID, field.TypeUUID, value)
	}
	if value, ok := duo.mutation.DishName(); ok {
		_spec.SetField(dish.FieldDishName, field.TypeString, value)
	}
	if value, ok := duo.mutation.Quantity(); ok {
		_spec.SetField(dish.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := duo.mutation.AddedQuantity(); ok {
		_spec.AddField(dish.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := duo.mutation.UpdateAt(); ok {
		_spec.SetField(dish.FieldUpdateAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.SetField(dish.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &Dish{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dish.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}