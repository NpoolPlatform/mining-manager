// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/mining-manager/pkg/db/ent/profitdetail"
	"github.com/NpoolPlatform/mining-manager/pkg/db/ent/profitgeneral"
	"github.com/NpoolPlatform/mining-manager/pkg/db/ent/profitunsold"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 3)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   profitdetail.Table,
			Columns: profitdetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profitdetail.FieldID,
			},
		},
		Type: "ProfitDetail",
		Fields: map[string]*sqlgraph.FieldSpec{
			profitdetail.FieldCreatedAt:   {Type: field.TypeUint32, Column: profitdetail.FieldCreatedAt},
			profitdetail.FieldUpdatedAt:   {Type: field.TypeUint32, Column: profitdetail.FieldUpdatedAt},
			profitdetail.FieldDeletedAt:   {Type: field.TypeUint32, Column: profitdetail.FieldDeletedAt},
			profitdetail.FieldGoodID:      {Type: field.TypeUUID, Column: profitdetail.FieldGoodID},
			profitdetail.FieldCoinTypeID:  {Type: field.TypeUUID, Column: profitdetail.FieldCoinTypeID},
			profitdetail.FieldAmount:      {Type: field.TypeFloat64, Column: profitdetail.FieldAmount},
			profitdetail.FieldBenefitDate: {Type: field.TypeUint32, Column: profitdetail.FieldBenefitDate},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   profitgeneral.Table,
			Columns: profitgeneral.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profitgeneral.FieldID,
			},
		},
		Type: "ProfitGeneral",
		Fields: map[string]*sqlgraph.FieldSpec{
			profitgeneral.FieldCreatedAt:  {Type: field.TypeUint32, Column: profitgeneral.FieldCreatedAt},
			profitgeneral.FieldUpdatedAt:  {Type: field.TypeUint32, Column: profitgeneral.FieldUpdatedAt},
			profitgeneral.FieldDeletedAt:  {Type: field.TypeUint32, Column: profitgeneral.FieldDeletedAt},
			profitgeneral.FieldGoodID:     {Type: field.TypeUUID, Column: profitgeneral.FieldGoodID},
			profitgeneral.FieldCoinTypeID: {Type: field.TypeUUID, Column: profitgeneral.FieldCoinTypeID},
			profitgeneral.FieldAmount:     {Type: field.TypeFloat64, Column: profitgeneral.FieldAmount},
			profitgeneral.FieldToPlatform: {Type: field.TypeFloat64, Column: profitgeneral.FieldToPlatform},
			profitgeneral.FieldToUser:     {Type: field.TypeFloat64, Column: profitgeneral.FieldToUser},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   profitunsold.Table,
			Columns: profitunsold.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profitunsold.FieldID,
			},
		},
		Type: "ProfitUnsold",
		Fields: map[string]*sqlgraph.FieldSpec{
			profitunsold.FieldCreatedAt:   {Type: field.TypeUint32, Column: profitunsold.FieldCreatedAt},
			profitunsold.FieldUpdatedAt:   {Type: field.TypeUint32, Column: profitunsold.FieldUpdatedAt},
			profitunsold.FieldDeletedAt:   {Type: field.TypeUint32, Column: profitunsold.FieldDeletedAt},
			profitunsold.FieldGoodID:      {Type: field.TypeUUID, Column: profitunsold.FieldGoodID},
			profitunsold.FieldCoinTypeID:  {Type: field.TypeUUID, Column: profitunsold.FieldCoinTypeID},
			profitunsold.FieldAmount:      {Type: field.TypeFloat64, Column: profitunsold.FieldAmount},
			profitunsold.FieldBenefitDate: {Type: field.TypeUint32, Column: profitunsold.FieldBenefitDate},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (pdq *ProfitDetailQuery) addPredicate(pred func(s *sql.Selector)) {
	pdq.predicates = append(pdq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ProfitDetailQuery builder.
func (pdq *ProfitDetailQuery) Filter() *ProfitDetailFilter {
	return &ProfitDetailFilter{pdq}
}

// addPredicate implements the predicateAdder interface.
func (m *ProfitDetailMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ProfitDetailMutation builder.
func (m *ProfitDetailMutation) Filter() *ProfitDetailFilter {
	return &ProfitDetailFilter{m}
}

// ProfitDetailFilter provides a generic filtering capability at runtime for ProfitDetailQuery.
type ProfitDetailFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *ProfitDetailFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *ProfitDetailFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(profitdetail.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ProfitDetailFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(profitdetail.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ProfitDetailFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(profitdetail.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ProfitDetailFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(profitdetail.FieldDeletedAt))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *ProfitDetailFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(profitdetail.FieldGoodID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *ProfitDetailFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(profitdetail.FieldCoinTypeID))
}

// WhereAmount applies the entql float64 predicate on the amount field.
func (f *ProfitDetailFilter) WhereAmount(p entql.Float64P) {
	f.Where(p.Field(profitdetail.FieldAmount))
}

// WhereBenefitDate applies the entql uint32 predicate on the benefit_date field.
func (f *ProfitDetailFilter) WhereBenefitDate(p entql.Uint32P) {
	f.Where(p.Field(profitdetail.FieldBenefitDate))
}

// addPredicate implements the predicateAdder interface.
func (pgq *ProfitGeneralQuery) addPredicate(pred func(s *sql.Selector)) {
	pgq.predicates = append(pgq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ProfitGeneralQuery builder.
func (pgq *ProfitGeneralQuery) Filter() *ProfitGeneralFilter {
	return &ProfitGeneralFilter{pgq}
}

// addPredicate implements the predicateAdder interface.
func (m *ProfitGeneralMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ProfitGeneralMutation builder.
func (m *ProfitGeneralMutation) Filter() *ProfitGeneralFilter {
	return &ProfitGeneralFilter{m}
}

// ProfitGeneralFilter provides a generic filtering capability at runtime for ProfitGeneralQuery.
type ProfitGeneralFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *ProfitGeneralFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *ProfitGeneralFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(profitgeneral.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ProfitGeneralFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(profitgeneral.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ProfitGeneralFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(profitgeneral.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ProfitGeneralFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(profitgeneral.FieldDeletedAt))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *ProfitGeneralFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(profitgeneral.FieldGoodID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *ProfitGeneralFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(profitgeneral.FieldCoinTypeID))
}

// WhereAmount applies the entql float64 predicate on the amount field.
func (f *ProfitGeneralFilter) WhereAmount(p entql.Float64P) {
	f.Where(p.Field(profitgeneral.FieldAmount))
}

// WhereToPlatform applies the entql float64 predicate on the to_platform field.
func (f *ProfitGeneralFilter) WhereToPlatform(p entql.Float64P) {
	f.Where(p.Field(profitgeneral.FieldToPlatform))
}

// WhereToUser applies the entql float64 predicate on the to_user field.
func (f *ProfitGeneralFilter) WhereToUser(p entql.Float64P) {
	f.Where(p.Field(profitgeneral.FieldToUser))
}

// addPredicate implements the predicateAdder interface.
func (puq *ProfitUnsoldQuery) addPredicate(pred func(s *sql.Selector)) {
	puq.predicates = append(puq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ProfitUnsoldQuery builder.
func (puq *ProfitUnsoldQuery) Filter() *ProfitUnsoldFilter {
	return &ProfitUnsoldFilter{puq}
}

// addPredicate implements the predicateAdder interface.
func (m *ProfitUnsoldMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ProfitUnsoldMutation builder.
func (m *ProfitUnsoldMutation) Filter() *ProfitUnsoldFilter {
	return &ProfitUnsoldFilter{m}
}

// ProfitUnsoldFilter provides a generic filtering capability at runtime for ProfitUnsoldQuery.
type ProfitUnsoldFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *ProfitUnsoldFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *ProfitUnsoldFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(profitunsold.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ProfitUnsoldFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(profitunsold.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ProfitUnsoldFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(profitunsold.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ProfitUnsoldFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(profitunsold.FieldDeletedAt))
}

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *ProfitUnsoldFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(profitunsold.FieldGoodID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *ProfitUnsoldFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(profitunsold.FieldCoinTypeID))
}

// WhereAmount applies the entql float64 predicate on the amount field.
func (f *ProfitUnsoldFilter) WhereAmount(p entql.Float64P) {
	f.Where(p.Field(profitunsold.FieldAmount))
}

// WhereBenefitDate applies the entql uint32 predicate on the benefit_date field.
func (f *ProfitUnsoldFilter) WhereBenefitDate(p entql.Uint32P) {
	f.Where(p.Field(profitunsold.FieldBenefitDate))
}
