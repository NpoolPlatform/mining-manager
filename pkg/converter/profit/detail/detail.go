package detail

import (
	npool "github.com/NpoolPlatform/message/npool/miningmgr/profit/detail"
	"github.com/NpoolPlatform/mining-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.ProfitDetail) *npool.ProfitDetail {
	return &npool.ProfitDetail{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		UserID:          row.UserID.String(),
		CoinTypeID:      row.CoinTypeID.String(),
		IOType:          npool.IOType(npool.IOType_value[row.IoType]),
		IOSubType:       npool.IOSubType(npool.IOSubType_value[row.IoSubType]),
		Amount:          row.Amount.String(),
		FromCoinTypeID:  row.FromCoinTypeID.String(),
		CoinUSDCurrency: row.CoinUsdCurrency.String(),
		IOExtra:         row.IoExtra,
		FromOldID:       row.FromOldID.String(),
	}
}

func Ent2GrpcMany(rows []*ent.ProfitDetail) []*npool.ProfitDetail {
	infos := []*npool.ProfitDetail{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
