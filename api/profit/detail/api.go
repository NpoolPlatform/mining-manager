package detail

import (
	"github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/profit/detail"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	detail.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	detail.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
