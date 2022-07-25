package unsold

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/mining-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/mining-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/mining-manager/pkg/tracer/profit/unsold"
	"github.com/shopspring/decimal"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/miningmgr/profit/unsold"
	"github.com/NpoolPlatform/mining-manager/pkg/db"
	"github.com/NpoolPlatform/mining-manager/pkg/db/ent"
	unsold "github.com/NpoolPlatform/mining-manager/pkg/db/ent/profitunsold"

	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.UnsoldReq) (*ent.Unsold, error) { //nolint
	var info *ent.Unsold
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.Debug().Unsold.Create()

		if in.ID != nil {
			c.SetID(uuid.MustParse(in.GetID()))
		}
		if in.AppID != nil {
			c.SetAppID(uuid.MustParse(in.GetAppID()))
		}
		if in.UserID != nil {
			c.SetUserID(uuid.MustParse(in.GetUserID()))
		}
		if in.CoinTypeID != nil {
			c.SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID()))
		}
		if in.IOType != nil {
			c.SetIoType(in.GetIOType().String())
		}
		if in.IOSubType != nil {
			c.SetIoSubType(in.GetIOSubType().String())
		}
		if in.Amount != nil {
			amount, err := decimal.NewFromString(in.GetAmount())
			if err != nil {
				return err
			}
			c.SetAmount(amount)
		}
		if in.FromCoinTypeID != nil {
			c.SetFromCoinTypeID(uuid.MustParse(in.GetFromCoinTypeID()))
		}
		if in.CoinUSDCurrency != nil {
			currency, err := decimal.NewFromString(in.GetCoinUSDCurrency())
			if err != nil {
				return err
			}
			c.SetCoinUsdCurrency(currency)
		}
		if in.IOExtra != nil {
			c.SetIoExtra(in.GetIOExtra())
		}
		if in.FromOldID != nil {
			c.SetFromOldID(uuid.MustParse(in.GetFromOldID()))
		}
		if in.CreatedAt != nil {
			c.SetCreatedAt(in.GetCreatedAt())
		}

		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.UnsoldReq) ([]*ent.Unsold, error) { //nolint
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.Unsold{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.UnsoldCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.Unsold.Create()
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.UserID != nil {
				bulk[i].SetUserID(uuid.MustParse(info.GetUserID()))
			}
			if info.CoinTypeID != nil {
				bulk[i].SetCoinTypeID(uuid.MustParse(info.GetCoinTypeID()))
			}
			if info.IOType != nil {
				bulk[i].SetIoType(info.GetIOType().String())
			}
			if info.IOSubType != nil {
				bulk[i].SetIoSubType(info.GetIOSubType().String())
			}
			if info.Amount != nil {
				amount, err := decimal.NewFromString(info.GetAmount())
				if err != nil {
					return err
				}
				bulk[i].SetAmount(amount)
			}
			if info.FromCoinTypeID != nil {
				bulk[i].SetCoinTypeID(uuid.MustParse(info.GetFromCoinTypeID()))
			}
			if info.CoinUSDCurrency != nil {
				currency, err := decimal.NewFromString(info.GetCoinUSDCurrency())
				if err != nil {
					return err
				}
				bulk[i].SetCoinUsdCurrency(currency)
			}
			if info.IOExtra != nil {
				bulk[i].SetIoExtra(info.GetIOExtra())
			}
			if info.FromOldID != nil {
				bulk[i].SetFromOldID(uuid.MustParse(info.GetFromOldID()))
			}
			if info.CreatedAt != nil {
				bulk[i].SetCreatedAt(info.GetCreatedAt())
			}
		}
		rows, err = tx.Unsold.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Unsold, error) {
	var info *ent.Unsold
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Unsold.Query().Where(unsold.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.UnsoldQuery, error) { //nolint
	stm := cli.Unsold.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(unsold.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid unsold field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(unsold.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid unsold field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(unsold.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid unsold field")
		}
	}
	if conds.CoinTypeID != nil {
		switch conds.GetCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(unsold.CoinTypeID(uuid.MustParse(conds.GetCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid unsold field")
		}
	}
	if conds.IOType != nil {
		switch conds.GetIOType().GetOp() {
		case cruder.EQ:
			stm.Where(unsold.IoType(npool.IOType(conds.GetIOType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid unsold field")
		}
	}
	if conds.IOSubType != nil {
		switch conds.GetIOSubType().GetOp() {
		case cruder.EQ:
			stm.Where(unsold.IoType(npool.IOSubType(conds.GetIOSubType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid unsold field")
		}
	}
	if conds.Amount != nil {
		amount, err := decimal.NewFromString(conds.GetAmount().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetAmount().GetOp() {
		case cruder.LT:
			stm.Where(unsold.AmountLT(amount))
		case cruder.GT:
			stm.Where(unsold.AmountGT(amount))
		case cruder.EQ:
			stm.Where(unsold.AmountEQ(amount))
		default:
			return nil, fmt.Errorf("invalid unsold field")
		}
	}
	if conds.FromCoinTypeID != nil {
		switch conds.GetFromCoinTypeID().GetOp() {
		case cruder.EQ:
			stm.Where(unsold.FromCoinTypeID(uuid.MustParse(conds.GetFromCoinTypeID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid unsold field")
		}
	}
	if conds.CoinUSDCurrency != nil {
		currency, err := decimal.NewFromString(conds.GetCoinUSDCurrency().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetCoinUSDCurrency().GetOp() {
		case cruder.LT:
			stm.Where(unsold.CoinUsdCurrencyLT(currency))
		case cruder.GT:
			stm.Where(unsold.CoinUsdCurrencyGT(currency))
		case cruder.EQ:
			stm.Where(unsold.CoinUsdCurrencyEQ(currency))
		default:
			return nil, fmt.Errorf("invalid unsold field")
		}
	}
	if conds.IOExtra != nil {
		switch conds.GetIOExtra().GetOp() {
		case cruder.LIKE:
			stm.Where(unsold.IoExtraContains(conds.GetIOExtra().GetValue()))
		default:
			return nil, fmt.Errorf("invalid unsold field")
		}
	}
	if conds.FromOldID != nil {
		switch conds.GetFromOldID().GetOp() {
		case cruder.EQ:
			stm.Where(unsold.FromOldID(uuid.MustParse(conds.GetFromOldID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid unsold field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Unsold, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.Unsold{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(unsold.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Unsold, error) {
	var info *ent.Unsold
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.Unsold.Query().Where(unsold.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.Unsold, error) {
	var info *ent.Unsold
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Unsold.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
