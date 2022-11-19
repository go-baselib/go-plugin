package hi

import (
	"context"

	bp "github.com/go-baselib/go-plugin"
	"github.com/go-baselib/go-plugin/pkg/server"

	"github.com/grpcprotocol/plugin"
)

const Name = "hi"

func init() {
	bp.Register(Name, server.NewGRPCPlugin(&Hi{}), "type.googleapis.com/proto.HiReq")
}

type Hi struct{}

func (h *Hi) Exec(ctx context.Context, in *plugin.ExecReq) (*plugin.ExecRsp, error) {
	var (
		req = &HiReq{}
		err error
	)
	if err = plugin.UnmarshalExecReq(in, req); err != nil {
		return nil, err
	}

	var (
		rsp = &HiRsp{Message: "Hi, " + req.GetName()}
		out = &plugin.ExecRsp{}
	)
	if out, err = plugin.MarshalExecRsp(rsp); err != nil {
		return nil, err
	}
	return out, nil
}
