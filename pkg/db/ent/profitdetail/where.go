// Code generated by entc, DO NOT EDIT.

package profitdetail

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/mining-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v decimal.Decimal) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// BenefitDate applies equality check predicate on the "benefit_date" field. It's identical to BenefitDateEQ.
func BenefitDate(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefitDate), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
	})
}

// GoodIDIsNil applies the IsNil predicate on the "good_id" field.
func GoodIDIsNil() predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGoodID)))
	})
}

// GoodIDNotNil applies the NotNil predicate on the "good_id" field.
func GoodIDNotNil() predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGoodID)))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinTypeID)))
	})
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinTypeID)))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v decimal.Decimal) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v decimal.Decimal) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...decimal.Decimal) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...decimal.Decimal) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v decimal.Decimal) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v decimal.Decimal) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v decimal.Decimal) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v decimal.Decimal) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// AmountIsNil applies the IsNil predicate on the "amount" field.
func AmountIsNil() predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAmount)))
	})
}

// AmountNotNil applies the NotNil predicate on the "amount" field.
func AmountNotNil() predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAmount)))
	})
}

// BenefitDateEQ applies the EQ predicate on the "benefit_date" field.
func BenefitDateEQ(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefitDate), v))
	})
}

// BenefitDateNEQ applies the NEQ predicate on the "benefit_date" field.
func BenefitDateNEQ(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBenefitDate), v))
	})
}

// BenefitDateIn applies the In predicate on the "benefit_date" field.
func BenefitDateIn(vs ...uint32) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBenefitDate), v...))
	})
}

// BenefitDateNotIn applies the NotIn predicate on the "benefit_date" field.
func BenefitDateNotIn(vs ...uint32) predicate.ProfitDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProfitDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBenefitDate), v...))
	})
}

// BenefitDateGT applies the GT predicate on the "benefit_date" field.
func BenefitDateGT(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBenefitDate), v))
	})
}

// BenefitDateGTE applies the GTE predicate on the "benefit_date" field.
func BenefitDateGTE(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBenefitDate), v))
	})
}

// BenefitDateLT applies the LT predicate on the "benefit_date" field.
func BenefitDateLT(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBenefitDate), v))
	})
}

// BenefitDateLTE applies the LTE predicate on the "benefit_date" field.
func BenefitDateLTE(v uint32) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBenefitDate), v))
	})
}

// BenefitDateIsNil applies the IsNil predicate on the "benefit_date" field.
func BenefitDateIsNil() predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBenefitDate)))
	})
}

// BenefitDateNotNil applies the NotNil predicate on the "benefit_date" field.
func BenefitDateNotNil() predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBenefitDate)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProfitDetail) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProfitDetail) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProfitDetail) predicate.ProfitDetail {
	return predicate.ProfitDetail(func(s *sql.Selector) {
		p(s.Not())
	})
}
