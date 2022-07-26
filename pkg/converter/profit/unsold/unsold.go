package unsold

import (
	npool "github.com/NpoolPlatform/message/npool/miningmgr/profit/unsold"
	"github.com/NpoolPlatform/mining-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.ProfitUnsold) *npool.Unsold {
	if row == nil {
		return nil
	}

	return &npool.Unsold{
		ID:          row.ID.String(),
		GoodID:      row.GoodID.String(),
		CoinTypeID:  row.CoinTypeID.String(),
		Amount:      row.Amount.String(),
		BenefitDate: row.BenefitDate,
		CreatedAt:   row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.ProfitUnsold) []*npool.Unsold {
	infos := []*npool.Unsold{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
