// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/profile"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ProfileCreate is the builder for creating a Profile entity.
type ProfileCreate struct {
	config
	mutation *ProfileMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNickname sets the "nickname" field.
func (pc *ProfileCreate) SetNickname(s string) *ProfileCreate {
	pc.mutation.SetNickname(s)
	return pc
}

// SetUUID sets the "uuid" field.
func (pc *ProfileCreate) SetUUID(u uuid.UUID) *ProfileCreate {
	pc.mutation.SetUUID(u)
	return pc
}

// SetIconURL sets the "icon_url" field.
func (pc *ProfileCreate) SetIconURL(s string) *ProfileCreate {
	pc.mutation.SetIconURL(s)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProfileCreate) SetCreatedAt(t time.Time) *ProfileCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableCreatedAt(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProfileCreate) SetUpdatedAt(t time.Time) *ProfileCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableUpdatedAt(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// Mutation returns the ProfileMutation object of the builder.
func (pc *ProfileCreate) Mutation() *ProfileMutation {
	return pc.mutation
}

// Save creates the Profile in the database.
func (pc *ProfileCreate) Save(ctx context.Context) (*Profile, error) {
	pc.defaults()
	return withHooks[*Profile, ProfileMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfileCreate) SaveX(ctx context.Context) *Profile {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProfileCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProfileCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProfileCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := profile.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := profile.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProfileCreate) check() error {
	if _, ok := pc.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`ent: missing required field "Profile.nickname"`)}
	}
	if _, ok := pc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Profile.uuid"`)}
	}
	if _, ok := pc.mutation.IconURL(); !ok {
		return &ValidationError{Name: "icon_url", err: errors.New(`ent: missing required field "Profile.icon_url"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Profile.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Profile.updated_at"`)}
	}
	return nil
}

func (pc *ProfileCreate) sqlSave(ctx context.Context) (*Profile, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProfileCreate) createSpec() (*Profile, *sqlgraph.CreateSpec) {
	var (
		_node = &Profile{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(profile.Table, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.Nickname(); ok {
		_spec.SetField(profile.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := pc.mutation.UUID(); ok {
		_spec.SetField(profile.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if value, ok := pc.mutation.IconURL(); ok {
		_spec.SetField(profile.FieldIconURL, field.TypeString, value)
		_node.IconURL = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(profile.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(profile.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Profile.Create().
//		SetNickname(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProfileUpsert) {
//			SetNickname(v+v).
//		}).
//		Exec(ctx)
func (pc *ProfileCreate) OnConflict(opts ...sql.ConflictOption) *ProfileUpsertOne {
	pc.conflict = opts
	return &ProfileUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Profile.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *ProfileCreate) OnConflictColumns(columns ...string) *ProfileUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &ProfileUpsertOne{
		create: pc,
	}
}

type (
	// ProfileUpsertOne is the builder for "upsert"-ing
	//  one Profile node.
	ProfileUpsertOne struct {
		create *ProfileCreate
	}

	// ProfileUpsert is the "OnConflict" setter.
	ProfileUpsert struct {
		*sql.UpdateSet
	}
)

// SetNickname sets the "nickname" field.
func (u *ProfileUpsert) SetNickname(v string) *ProfileUpsert {
	u.Set(profile.FieldNickname, v)
	return u
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *ProfileUpsert) UpdateNickname() *ProfileUpsert {
	u.SetExcluded(profile.FieldNickname)
	return u
}

// SetIconURL sets the "icon_url" field.
func (u *ProfileUpsert) SetIconURL(v string) *ProfileUpsert {
	u.Set(profile.FieldIconURL, v)
	return u
}

// UpdateIconURL sets the "icon_url" field to the value that was provided on create.
func (u *ProfileUpsert) UpdateIconURL() *ProfileUpsert {
	u.SetExcluded(profile.FieldIconURL)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ProfileUpsert) SetCreatedAt(v time.Time) *ProfileUpsert {
	u.Set(profile.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ProfileUpsert) UpdateCreatedAt() *ProfileUpsert {
	u.SetExcluded(profile.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProfileUpsert) SetUpdatedAt(v time.Time) *ProfileUpsert {
	u.Set(profile.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProfileUpsert) UpdateUpdatedAt() *ProfileUpsert {
	u.SetExcluded(profile.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Profile.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProfileUpsertOne) UpdateNewValues() *ProfileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.UUID(); exists {
			s.SetIgnore(profile.FieldUUID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Profile.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ProfileUpsertOne) Ignore() *ProfileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProfileUpsertOne) DoNothing() *ProfileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProfileCreate.OnConflict
// documentation for more info.
func (u *ProfileUpsertOne) Update(set func(*ProfileUpsert)) *ProfileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProfileUpsert{UpdateSet: update})
	}))
	return u
}

// SetNickname sets the "nickname" field.
func (u *ProfileUpsertOne) SetNickname(v string) *ProfileUpsertOne {
	return u.Update(func(s *ProfileUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *ProfileUpsertOne) UpdateNickname() *ProfileUpsertOne {
	return u.Update(func(s *ProfileUpsert) {
		s.UpdateNickname()
	})
}

// SetIconURL sets the "icon_url" field.
func (u *ProfileUpsertOne) SetIconURL(v string) *ProfileUpsertOne {
	return u.Update(func(s *ProfileUpsert) {
		s.SetIconURL(v)
	})
}

// UpdateIconURL sets the "icon_url" field to the value that was provided on create.
func (u *ProfileUpsertOne) UpdateIconURL() *ProfileUpsertOne {
	return u.Update(func(s *ProfileUpsert) {
		s.UpdateIconURL()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ProfileUpsertOne) SetCreatedAt(v time.Time) *ProfileUpsertOne {
	return u.Update(func(s *ProfileUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ProfileUpsertOne) UpdateCreatedAt() *ProfileUpsertOne {
	return u.Update(func(s *ProfileUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProfileUpsertOne) SetUpdatedAt(v time.Time) *ProfileUpsertOne {
	return u.Update(func(s *ProfileUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProfileUpsertOne) UpdateUpdatedAt() *ProfileUpsertOne {
	return u.Update(func(s *ProfileUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ProfileUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProfileCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProfileUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProfileUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProfileUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProfileCreateBulk is the builder for creating many Profile entities in bulk.
type ProfileCreateBulk struct {
	config
	builders []*ProfileCreate
	conflict []sql.ConflictOption
}

// Save creates the Profile entities in the database.
func (pcb *ProfileCreateBulk) Save(ctx context.Context) ([]*Profile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Profile, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProfileCreateBulk) SaveX(ctx context.Context) []*Profile {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProfileCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Profile.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProfileUpsert) {
//			SetNickname(v+v).
//		}).
//		Exec(ctx)
func (pcb *ProfileCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProfileUpsertBulk {
	pcb.conflict = opts
	return &ProfileUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Profile.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *ProfileCreateBulk) OnConflictColumns(columns ...string) *ProfileUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &ProfileUpsertBulk{
		create: pcb,
	}
}

// ProfileUpsertBulk is the builder for "upsert"-ing
// a bulk of Profile nodes.
type ProfileUpsertBulk struct {
	create *ProfileCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Profile.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProfileUpsertBulk) UpdateNewValues() *ProfileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.UUID(); exists {
				s.SetIgnore(profile.FieldUUID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Profile.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ProfileUpsertBulk) Ignore() *ProfileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProfileUpsertBulk) DoNothing() *ProfileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProfileCreateBulk.OnConflict
// documentation for more info.
func (u *ProfileUpsertBulk) Update(set func(*ProfileUpsert)) *ProfileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProfileUpsert{UpdateSet: update})
	}))
	return u
}

// SetNickname sets the "nickname" field.
func (u *ProfileUpsertBulk) SetNickname(v string) *ProfileUpsertBulk {
	return u.Update(func(s *ProfileUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *ProfileUpsertBulk) UpdateNickname() *ProfileUpsertBulk {
	return u.Update(func(s *ProfileUpsert) {
		s.UpdateNickname()
	})
}

// SetIconURL sets the "icon_url" field.
func (u *ProfileUpsertBulk) SetIconURL(v string) *ProfileUpsertBulk {
	return u.Update(func(s *ProfileUpsert) {
		s.SetIconURL(v)
	})
}

// UpdateIconURL sets the "icon_url" field to the value that was provided on create.
func (u *ProfileUpsertBulk) UpdateIconURL() *ProfileUpsertBulk {
	return u.Update(func(s *ProfileUpsert) {
		s.UpdateIconURL()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ProfileUpsertBulk) SetCreatedAt(v time.Time) *ProfileUpsertBulk {
	return u.Update(func(s *ProfileUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ProfileUpsertBulk) UpdateCreatedAt() *ProfileUpsertBulk {
	return u.Update(func(s *ProfileUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProfileUpsertBulk) SetUpdatedAt(v time.Time) *ProfileUpsertBulk {
	return u.Update(func(s *ProfileUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProfileUpsertBulk) UpdateUpdatedAt() *ProfileUpsertBulk {
	return u.Update(func(s *ProfileUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ProfileUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProfileCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProfileCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProfileUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
