package detail

import (
	"github.com/NpoolPlatform/message/npool/miningmgr/profit/detail"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	detail.UnimplementedProfitDetailServer
}

func Register(server grpc.ServiceRegistrar) {
	detail.RegisterProfitDetailServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
