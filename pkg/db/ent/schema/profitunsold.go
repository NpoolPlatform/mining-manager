//nolint:nolintlint,dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/mining-manager/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ProfitUnsold holds the schema definition for the ProfitUnsold entity.
type ProfitUnsold struct {
	ent.Schema
}

func (ProfitUnsold) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the ProfitUnsold.
func (ProfitUnsold) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Float("amount").
			GoType(decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37, 18)",
			}).
			Optional(),
		field.
			Uint32("benefit_date").
			Optional().
			Default(0),
	}
}

// Edges of the ProfitUnsold.
func (ProfitUnsold) Edges() []ent.Edge {
	return nil
}
