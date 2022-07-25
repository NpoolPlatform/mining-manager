// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/mining-manager/pkg/db/ent/profitdetail"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ProfitDetailCreate is the builder for creating a ProfitDetail entity.
type ProfitDetailCreate struct {
	config
	mutation *ProfitDetailMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (pdc *ProfitDetailCreate) SetCreatedAt(u uint32) *ProfitDetailCreate {
	pdc.mutation.SetCreatedAt(u)
	return pdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pdc *ProfitDetailCreate) SetNillableCreatedAt(u *uint32) *ProfitDetailCreate {
	if u != nil {
		pdc.SetCreatedAt(*u)
	}
	return pdc
}

// SetUpdatedAt sets the "updated_at" field.
func (pdc *ProfitDetailCreate) SetUpdatedAt(u uint32) *ProfitDetailCreate {
	pdc.mutation.SetUpdatedAt(u)
	return pdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pdc *ProfitDetailCreate) SetNillableUpdatedAt(u *uint32) *ProfitDetailCreate {
	if u != nil {
		pdc.SetUpdatedAt(*u)
	}
	return pdc
}

// SetDeletedAt sets the "deleted_at" field.
func (pdc *ProfitDetailCreate) SetDeletedAt(u uint32) *ProfitDetailCreate {
	pdc.mutation.SetDeletedAt(u)
	return pdc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pdc *ProfitDetailCreate) SetNillableDeletedAt(u *uint32) *ProfitDetailCreate {
	if u != nil {
		pdc.SetDeletedAt(*u)
	}
	return pdc
}

// SetGoodID sets the "good_id" field.
func (pdc *ProfitDetailCreate) SetGoodID(u uuid.UUID) *ProfitDetailCreate {
	pdc.mutation.SetGoodID(u)
	return pdc
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (pdc *ProfitDetailCreate) SetNillableGoodID(u *uuid.UUID) *ProfitDetailCreate {
	if u != nil {
		pdc.SetGoodID(*u)
	}
	return pdc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (pdc *ProfitDetailCreate) SetCoinTypeID(u uuid.UUID) *ProfitDetailCreate {
	pdc.mutation.SetCoinTypeID(u)
	return pdc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (pdc *ProfitDetailCreate) SetNillableCoinTypeID(u *uuid.UUID) *ProfitDetailCreate {
	if u != nil {
		pdc.SetCoinTypeID(*u)
	}
	return pdc
}

// SetAmount sets the "amount" field.
func (pdc *ProfitDetailCreate) SetAmount(d decimal.Decimal) *ProfitDetailCreate {
	pdc.mutation.SetAmount(d)
	return pdc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pdc *ProfitDetailCreate) SetNillableAmount(d *decimal.Decimal) *ProfitDetailCreate {
	if d != nil {
		pdc.SetAmount(*d)
	}
	return pdc
}

// SetBenefitDate sets the "benefit_date" field.
func (pdc *ProfitDetailCreate) SetBenefitDate(u uint32) *ProfitDetailCreate {
	pdc.mutation.SetBenefitDate(u)
	return pdc
}

// SetNillableBenefitDate sets the "benefit_date" field if the given value is not nil.
func (pdc *ProfitDetailCreate) SetNillableBenefitDate(u *uint32) *ProfitDetailCreate {
	if u != nil {
		pdc.SetBenefitDate(*u)
	}
	return pdc
}

// SetID sets the "id" field.
func (pdc *ProfitDetailCreate) SetID(u uuid.UUID) *ProfitDetailCreate {
	pdc.mutation.SetID(u)
	return pdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pdc *ProfitDetailCreate) SetNillableID(u *uuid.UUID) *ProfitDetailCreate {
	if u != nil {
		pdc.SetID(*u)
	}
	return pdc
}

// Mutation returns the ProfitDetailMutation object of the builder.
func (pdc *ProfitDetailCreate) Mutation() *ProfitDetailMutation {
	return pdc.mutation
}

// Save creates the ProfitDetail in the database.
func (pdc *ProfitDetailCreate) Save(ctx context.Context) (*ProfitDetail, error) {
	var (
		err  error
		node *ProfitDetail
	)
	if err := pdc.defaults(); err != nil {
		return nil, err
	}
	if len(pdc.hooks) == 0 {
		if err = pdc.check(); err != nil {
			return nil, err
		}
		node, err = pdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfitDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pdc.check(); err != nil {
				return nil, err
			}
			pdc.mutation = mutation
			if node, err = pdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pdc.hooks) - 1; i >= 0; i-- {
			if pdc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pdc *ProfitDetailCreate) SaveX(ctx context.Context) *ProfitDetail {
	v, err := pdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pdc *ProfitDetailCreate) Exec(ctx context.Context) error {
	_, err := pdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdc *ProfitDetailCreate) ExecX(ctx context.Context) {
	if err := pdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pdc *ProfitDetailCreate) defaults() error {
	if _, ok := pdc.mutation.CreatedAt(); !ok {
		if profitdetail.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized profitdetail.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := profitdetail.DefaultCreatedAt()
		pdc.mutation.SetCreatedAt(v)
	}
	if _, ok := pdc.mutation.UpdatedAt(); !ok {
		if profitdetail.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized profitdetail.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := profitdetail.DefaultUpdatedAt()
		pdc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pdc.mutation.DeletedAt(); !ok {
		if profitdetail.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized profitdetail.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := profitdetail.DefaultDeletedAt()
		pdc.mutation.SetDeletedAt(v)
	}
	if _, ok := pdc.mutation.GoodID(); !ok {
		if profitdetail.DefaultGoodID == nil {
			return fmt.Errorf("ent: uninitialized profitdetail.DefaultGoodID (forgotten import ent/runtime?)")
		}
		v := profitdetail.DefaultGoodID()
		pdc.mutation.SetGoodID(v)
	}
	if _, ok := pdc.mutation.CoinTypeID(); !ok {
		if profitdetail.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized profitdetail.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := profitdetail.DefaultCoinTypeID()
		pdc.mutation.SetCoinTypeID(v)
	}
	if _, ok := pdc.mutation.BenefitDate(); !ok {
		v := profitdetail.DefaultBenefitDate
		pdc.mutation.SetBenefitDate(v)
	}
	if _, ok := pdc.mutation.ID(); !ok {
		if profitdetail.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized profitdetail.DefaultID (forgotten import ent/runtime?)")
		}
		v := profitdetail.DefaultID()
		pdc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pdc *ProfitDetailCreate) check() error {
	if _, ok := pdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProfitDetail.created_at"`)}
	}
	if _, ok := pdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProfitDetail.updated_at"`)}
	}
	if _, ok := pdc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "ProfitDetail.deleted_at"`)}
	}
	return nil
}

func (pdc *ProfitDetailCreate) sqlSave(ctx context.Context) (*ProfitDetail, error) {
	_node, _spec := pdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (pdc *ProfitDetailCreate) createSpec() (*ProfitDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &ProfitDetail{config: pdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: profitdetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profitdetail.FieldID,
			},
		}
	)
	_spec.OnConflict = pdc.conflict
	if id, ok := pdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pdc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: profitdetail.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pdc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: profitdetail.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pdc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: profitdetail.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := pdc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: profitdetail.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := pdc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: profitdetail.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := pdc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: profitdetail.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := pdc.mutation.BenefitDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: profitdetail.FieldBenefitDate,
		})
		_node.BenefitDate = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProfitDetail.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProfitDetailUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (pdc *ProfitDetailCreate) OnConflict(opts ...sql.ConflictOption) *ProfitDetailUpsertOne {
	pdc.conflict = opts
	return &ProfitDetailUpsertOne{
		create: pdc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProfitDetail.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pdc *ProfitDetailCreate) OnConflictColumns(columns ...string) *ProfitDetailUpsertOne {
	pdc.conflict = append(pdc.conflict, sql.ConflictColumns(columns...))
	return &ProfitDetailUpsertOne{
		create: pdc,
	}
}

type (
	// ProfitDetailUpsertOne is the builder for "upsert"-ing
	//  one ProfitDetail node.
	ProfitDetailUpsertOne struct {
		create *ProfitDetailCreate
	}

	// ProfitDetailUpsert is the "OnConflict" setter.
	ProfitDetailUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *ProfitDetailUpsert) SetCreatedAt(v uint32) *ProfitDetailUpsert {
	u.Set(profitdetail.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ProfitDetailUpsert) UpdateCreatedAt() *ProfitDetailUpsert {
	u.SetExcluded(profitdetail.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ProfitDetailUpsert) AddCreatedAt(v uint32) *ProfitDetailUpsert {
	u.Add(profitdetail.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProfitDetailUpsert) SetUpdatedAt(v uint32) *ProfitDetailUpsert {
	u.Set(profitdetail.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProfitDetailUpsert) UpdateUpdatedAt() *ProfitDetailUpsert {
	u.SetExcluded(profitdetail.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ProfitDetailUpsert) AddUpdatedAt(v uint32) *ProfitDetailUpsert {
	u.Add(profitdetail.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ProfitDetailUpsert) SetDeletedAt(v uint32) *ProfitDetailUpsert {
	u.Set(profitdetail.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ProfitDetailUpsert) UpdateDeletedAt() *ProfitDetailUpsert {
	u.SetExcluded(profitdetail.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ProfitDetailUpsert) AddDeletedAt(v uint32) *ProfitDetailUpsert {
	u.Add(profitdetail.FieldDeletedAt, v)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *ProfitDetailUpsert) SetGoodID(v uuid.UUID) *ProfitDetailUpsert {
	u.Set(profitdetail.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *ProfitDetailUpsert) UpdateGoodID() *ProfitDetailUpsert {
	u.SetExcluded(profitdetail.FieldGoodID)
	return u
}

// ClearGoodID clears the value of the "good_id" field.
func (u *ProfitDetailUpsert) ClearGoodID() *ProfitDetailUpsert {
	u.SetNull(profitdetail.FieldGoodID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *ProfitDetailUpsert) SetCoinTypeID(v uuid.UUID) *ProfitDetailUpsert {
	u.Set(profitdetail.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *ProfitDetailUpsert) UpdateCoinTypeID() *ProfitDetailUpsert {
	u.SetExcluded(profitdetail.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *ProfitDetailUpsert) ClearCoinTypeID() *ProfitDetailUpsert {
	u.SetNull(profitdetail.FieldCoinTypeID)
	return u
}

// SetAmount sets the "amount" field.
func (u *ProfitDetailUpsert) SetAmount(v decimal.Decimal) *ProfitDetailUpsert {
	u.Set(profitdetail.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *ProfitDetailUpsert) UpdateAmount() *ProfitDetailUpsert {
	u.SetExcluded(profitdetail.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *ProfitDetailUpsert) AddAmount(v decimal.Decimal) *ProfitDetailUpsert {
	u.Add(profitdetail.FieldAmount, v)
	return u
}

// ClearAmount clears the value of the "amount" field.
func (u *ProfitDetailUpsert) ClearAmount() *ProfitDetailUpsert {
	u.SetNull(profitdetail.FieldAmount)
	return u
}

// SetBenefitDate sets the "benefit_date" field.
func (u *ProfitDetailUpsert) SetBenefitDate(v uint32) *ProfitDetailUpsert {
	u.Set(profitdetail.FieldBenefitDate, v)
	return u
}

// UpdateBenefitDate sets the "benefit_date" field to the value that was provided on create.
func (u *ProfitDetailUpsert) UpdateBenefitDate() *ProfitDetailUpsert {
	u.SetExcluded(profitdetail.FieldBenefitDate)
	return u
}

// AddBenefitDate adds v to the "benefit_date" field.
func (u *ProfitDetailUpsert) AddBenefitDate(v uint32) *ProfitDetailUpsert {
	u.Add(profitdetail.FieldBenefitDate, v)
	return u
}

// ClearBenefitDate clears the value of the "benefit_date" field.
func (u *ProfitDetailUpsert) ClearBenefitDate() *ProfitDetailUpsert {
	u.SetNull(profitdetail.FieldBenefitDate)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ProfitDetail.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(profitdetail.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ProfitDetailUpsertOne) UpdateNewValues() *ProfitDetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(profitdetail.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ProfitDetail.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ProfitDetailUpsertOne) Ignore() *ProfitDetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProfitDetailUpsertOne) DoNothing() *ProfitDetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProfitDetailCreate.OnConflict
// documentation for more info.
func (u *ProfitDetailUpsertOne) Update(set func(*ProfitDetailUpsert)) *ProfitDetailUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProfitDetailUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ProfitDetailUpsertOne) SetCreatedAt(v uint32) *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ProfitDetailUpsertOne) AddCreatedAt(v uint32) *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ProfitDetailUpsertOne) UpdateCreatedAt() *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProfitDetailUpsertOne) SetUpdatedAt(v uint32) *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ProfitDetailUpsertOne) AddUpdatedAt(v uint32) *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProfitDetailUpsertOne) UpdateUpdatedAt() *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ProfitDetailUpsertOne) SetDeletedAt(v uint32) *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ProfitDetailUpsertOne) AddDeletedAt(v uint32) *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ProfitDetailUpsertOne) UpdateDeletedAt() *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetGoodID sets the "good_id" field.
func (u *ProfitDetailUpsertOne) SetGoodID(v uuid.UUID) *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *ProfitDetailUpsertOne) UpdateGoodID() *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *ProfitDetailUpsertOne) ClearGoodID() *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.ClearGoodID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *ProfitDetailUpsertOne) SetCoinTypeID(v uuid.UUID) *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *ProfitDetailUpsertOne) UpdateCoinTypeID() *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *ProfitDetailUpsertOne) ClearCoinTypeID() *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetAmount sets the "amount" field.
func (u *ProfitDetailUpsertOne) SetAmount(v decimal.Decimal) *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *ProfitDetailUpsertOne) AddAmount(v decimal.Decimal) *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *ProfitDetailUpsertOne) UpdateAmount() *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateAmount()
	})
}

// ClearAmount clears the value of the "amount" field.
func (u *ProfitDetailUpsertOne) ClearAmount() *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.ClearAmount()
	})
}

// SetBenefitDate sets the "benefit_date" field.
func (u *ProfitDetailUpsertOne) SetBenefitDate(v uint32) *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetBenefitDate(v)
	})
}

// AddBenefitDate adds v to the "benefit_date" field.
func (u *ProfitDetailUpsertOne) AddBenefitDate(v uint32) *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.AddBenefitDate(v)
	})
}

// UpdateBenefitDate sets the "benefit_date" field to the value that was provided on create.
func (u *ProfitDetailUpsertOne) UpdateBenefitDate() *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateBenefitDate()
	})
}

// ClearBenefitDate clears the value of the "benefit_date" field.
func (u *ProfitDetailUpsertOne) ClearBenefitDate() *ProfitDetailUpsertOne {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.ClearBenefitDate()
	})
}

// Exec executes the query.
func (u *ProfitDetailUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProfitDetailCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProfitDetailUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProfitDetailUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ProfitDetailUpsertOne.ID is not supported by MySQL driver. Use ProfitDetailUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProfitDetailUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProfitDetailCreateBulk is the builder for creating many ProfitDetail entities in bulk.
type ProfitDetailCreateBulk struct {
	config
	builders []*ProfitDetailCreate
	conflict []sql.ConflictOption
}

// Save creates the ProfitDetail entities in the database.
func (pdcb *ProfitDetailCreateBulk) Save(ctx context.Context) ([]*ProfitDetail, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pdcb.builders))
	nodes := make([]*ProfitDetail, len(pdcb.builders))
	mutators := make([]Mutator, len(pdcb.builders))
	for i := range pdcb.builders {
		func(i int, root context.Context) {
			builder := pdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfitDetailMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pdcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, pdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pdcb *ProfitDetailCreateBulk) SaveX(ctx context.Context) []*ProfitDetail {
	v, err := pdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pdcb *ProfitDetailCreateBulk) Exec(ctx context.Context) error {
	_, err := pdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdcb *ProfitDetailCreateBulk) ExecX(ctx context.Context) {
	if err := pdcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProfitDetail.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProfitDetailUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (pdcb *ProfitDetailCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProfitDetailUpsertBulk {
	pdcb.conflict = opts
	return &ProfitDetailUpsertBulk{
		create: pdcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProfitDetail.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pdcb *ProfitDetailCreateBulk) OnConflictColumns(columns ...string) *ProfitDetailUpsertBulk {
	pdcb.conflict = append(pdcb.conflict, sql.ConflictColumns(columns...))
	return &ProfitDetailUpsertBulk{
		create: pdcb,
	}
}

// ProfitDetailUpsertBulk is the builder for "upsert"-ing
// a bulk of ProfitDetail nodes.
type ProfitDetailUpsertBulk struct {
	create *ProfitDetailCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ProfitDetail.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(profitdetail.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ProfitDetailUpsertBulk) UpdateNewValues() *ProfitDetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(profitdetail.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ProfitDetail.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ProfitDetailUpsertBulk) Ignore() *ProfitDetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProfitDetailUpsertBulk) DoNothing() *ProfitDetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProfitDetailCreateBulk.OnConflict
// documentation for more info.
func (u *ProfitDetailUpsertBulk) Update(set func(*ProfitDetailUpsert)) *ProfitDetailUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProfitDetailUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ProfitDetailUpsertBulk) SetCreatedAt(v uint32) *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ProfitDetailUpsertBulk) AddCreatedAt(v uint32) *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ProfitDetailUpsertBulk) UpdateCreatedAt() *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProfitDetailUpsertBulk) SetUpdatedAt(v uint32) *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ProfitDetailUpsertBulk) AddUpdatedAt(v uint32) *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProfitDetailUpsertBulk) UpdateUpdatedAt() *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ProfitDetailUpsertBulk) SetDeletedAt(v uint32) *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ProfitDetailUpsertBulk) AddDeletedAt(v uint32) *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ProfitDetailUpsertBulk) UpdateDeletedAt() *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetGoodID sets the "good_id" field.
func (u *ProfitDetailUpsertBulk) SetGoodID(v uuid.UUID) *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *ProfitDetailUpsertBulk) UpdateGoodID() *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *ProfitDetailUpsertBulk) ClearGoodID() *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.ClearGoodID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *ProfitDetailUpsertBulk) SetCoinTypeID(v uuid.UUID) *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *ProfitDetailUpsertBulk) UpdateCoinTypeID() *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *ProfitDetailUpsertBulk) ClearCoinTypeID() *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetAmount sets the "amount" field.
func (u *ProfitDetailUpsertBulk) SetAmount(v decimal.Decimal) *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *ProfitDetailUpsertBulk) AddAmount(v decimal.Decimal) *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *ProfitDetailUpsertBulk) UpdateAmount() *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateAmount()
	})
}

// ClearAmount clears the value of the "amount" field.
func (u *ProfitDetailUpsertBulk) ClearAmount() *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.ClearAmount()
	})
}

// SetBenefitDate sets the "benefit_date" field.
func (u *ProfitDetailUpsertBulk) SetBenefitDate(v uint32) *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.SetBenefitDate(v)
	})
}

// AddBenefitDate adds v to the "benefit_date" field.
func (u *ProfitDetailUpsertBulk) AddBenefitDate(v uint32) *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.AddBenefitDate(v)
	})
}

// UpdateBenefitDate sets the "benefit_date" field to the value that was provided on create.
func (u *ProfitDetailUpsertBulk) UpdateBenefitDate() *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.UpdateBenefitDate()
	})
}

// ClearBenefitDate clears the value of the "benefit_date" field.
func (u *ProfitDetailUpsertBulk) ClearBenefitDate() *ProfitDetailUpsertBulk {
	return u.Update(func(s *ProfitDetailUpsert) {
		s.ClearBenefitDate()
	})
}

// Exec executes the query.
func (u *ProfitDetailUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProfitDetailCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProfitDetailCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProfitDetailUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}