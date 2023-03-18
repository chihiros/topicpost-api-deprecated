// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/predicate"
	"app/ent/prefecture"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PrefectureUpdate is the builder for updating Prefecture entities.
type PrefectureUpdate struct {
	config
	hooks    []Hook
	mutation *PrefectureMutation
}

// Where appends a list predicates to the PrefectureUpdate builder.
func (pu *PrefectureUpdate) Where(ps ...predicate.Prefecture) *PrefectureUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PrefectureUpdate) SetCreatedAt(t time.Time) *PrefectureUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PrefectureUpdate) SetNillableCreatedAt(t *time.Time) *PrefectureUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PrefectureUpdate) SetUpdatedAt(t time.Time) *PrefectureUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *PrefectureUpdate) SetNillableUpdatedAt(t *time.Time) *PrefectureUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// Mutation returns the PrefectureMutation object of the builder.
func (pu *PrefectureUpdate) Mutation() *PrefectureMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PrefectureUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PrefectureMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PrefectureUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PrefectureUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PrefectureUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PrefectureUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(prefecture.Table, prefecture.Columns, sqlgraph.NewFieldSpec(prefecture.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(prefecture.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(prefecture.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prefecture.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PrefectureUpdateOne is the builder for updating a single Prefecture entity.
type PrefectureUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PrefectureMutation
}

// SetCreatedAt sets the "created_at" field.
func (puo *PrefectureUpdateOne) SetCreatedAt(t time.Time) *PrefectureUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PrefectureUpdateOne) SetNillableCreatedAt(t *time.Time) *PrefectureUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PrefectureUpdateOne) SetUpdatedAt(t time.Time) *PrefectureUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *PrefectureUpdateOne) SetNillableUpdatedAt(t *time.Time) *PrefectureUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// Mutation returns the PrefectureMutation object of the builder.
func (puo *PrefectureUpdateOne) Mutation() *PrefectureMutation {
	return puo.mutation
}

// Where appends a list predicates to the PrefectureUpdate builder.
func (puo *PrefectureUpdateOne) Where(ps ...predicate.Prefecture) *PrefectureUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PrefectureUpdateOne) Select(field string, fields ...string) *PrefectureUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Prefecture entity.
func (puo *PrefectureUpdateOne) Save(ctx context.Context) (*Prefecture, error) {
	return withHooks[*Prefecture, PrefectureMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PrefectureUpdateOne) SaveX(ctx context.Context) *Prefecture {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PrefectureUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PrefectureUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PrefectureUpdateOne) sqlSave(ctx context.Context) (_node *Prefecture, err error) {
	_spec := sqlgraph.NewUpdateSpec(prefecture.Table, prefecture.Columns, sqlgraph.NewFieldSpec(prefecture.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Prefecture.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prefecture.FieldID)
		for _, f := range fields {
			if !prefecture.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != prefecture.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(prefecture.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(prefecture.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Prefecture{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prefecture.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
