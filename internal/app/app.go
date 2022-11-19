package app

import (
	"fmt"
	"net"

	"github.com/go-baselib/go-plugin/config"
	"github.com/go-baselib/go-plugin/internal/controller/grpc"
	"github.com/go-baselib/go-plugin/internal/usecase"
	"github.com/go-baselib/go-plugin/internal/usecase/api"
	"github.com/go-baselib/go-plugin/internal/usecase/repo"
)

func Run(cfg *config.Config) {
	var (
		pluginUseCase = usecase.NewPlugin(repo.NewPlugin(), api.NewPlugin())
		server        = grpc.NewServer(pluginUseCase)
		lis           net.Listener
		err           error
	)

	if lis, err = net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)); err != nil {
		panic(err)
	}

	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
