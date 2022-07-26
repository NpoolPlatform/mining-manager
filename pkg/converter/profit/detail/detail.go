package detail

import (
	npool "github.com/NpoolPlatform/message/npool/miningmgr/profit/detail"
	"github.com/NpoolPlatform/mining-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.ProfitDetail) *npool.Detail {
	if row == nil {
		return nil
	}

	return &npool.Detail{
		ID:          row.ID.String(),
		GoodID:      row.GoodID.String(),
		CoinTypeID:  row.CoinTypeID.String(),
		Amount:      row.Amount.String(),
		BenefitDate: row.BenefitDate,
		CreatedAt:   row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.ProfitDetail) []*npool.Detail {
	infos := []*npool.Detail{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
