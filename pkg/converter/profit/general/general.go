package general

import (
	npool "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/profit/general"
	"github.com/NpoolPlatform/mining-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.ProfitGeneral) *npool.General {
	if row == nil {
		return nil
	}

	return &npool.General{
		ID:                    row.ID.String(),
		GoodID:                row.GoodID.String(),
		CoinTypeID:            row.CoinTypeID.String(),
		Amount:                row.Amount.String(),
		ToPlatform:            row.ToPlatform.String(),
		ToUser:                row.ToUser.String(),
		TransferredToPlatform: row.ToPlatform.String(),
		TransferredToUser:     row.ToUser.String(),
	}
}

func Ent2GrpcMany(rows []*ent.ProfitGeneral) []*npool.General {
	infos := []*npool.General{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
