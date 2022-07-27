package unsold

import (
	"github.com/NpoolPlatform/message/npool/mining/mgr/v1/profit/unsold"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	unsold.UnimplementedProfitUnsoldServer
}

func Register(server grpc.ServiceRegistrar) {
	unsold.RegisterProfitUnsoldServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
