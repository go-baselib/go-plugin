package app

import (
	"fmt"
	"net"

	"github.com/go-baselib/go-plugin/config"
	"github.com/go-baselib/go-plugin/internal/controller/grpc"
	"github.com/go-baselib/go-plugin/internal/usecase"
	"github.com/go-baselib/go-plugin/internal/usecase/api"
	"github.com/go-baselib/go-plugin/internal/usecase/repo"
	"github.com/go-baselib/go-plugin/pkg/db"
)

func Run(cfg *config.Config) {
	var gDB, err = db.New(&cfg.DB)
	if err != nil {
		panic(err)
	}

	var (
		pluginUseCase = usecase.NewPlugin(repo.NewPlugin(gDB), api.NewPlugin())
		server        = grpc.NewServer(pluginUseCase)
		lis           net.Listener
	)

	if lis, err = net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)); err != nil {
		panic(err)
	}

	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
