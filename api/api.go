package api

import (
	"context"

	"github.com/NpoolPlatform/message/npool/miningmgr"

	"github.com/NpoolPlatform/mining-manager/api/profit/detail"
	"github.com/NpoolPlatform/mining-manager/api/profit/general"
	"github.com/NpoolPlatform/mining-manager/api/profit/unsold"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	miningmgr.UnimplementedMiningManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	miningmgr.RegisterMiningManagerServer(server, &Server{})
	general.Register(server)
	detail.Register(server)
	unsold.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := miningmgr.RegisterMiningManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := general.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := detail.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := unsold.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
