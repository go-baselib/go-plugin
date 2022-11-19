package main

import (
	"context"
	"fmt"

	"github.com/go-baselib/go-plugin/pkg/hi"

	"github.com/grpcprotocol/plugin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	var conn, err = grpc.Dial("127.0.0.1:9999", opts...)
	if err != nil {
		panic(err)
	}

	var (
		cli = plugin.NewPluginClient(conn)
		req = &hi.HiReq{Name: "jeson"}
		in  *plugin.ExecReq
	)
	if in, err = plugin.MarshalExecReq(req); err != nil {
		panic(err)
	}

	var (
		out *plugin.ExecRsp
		rsp = &hi.HiRsp{}
	)
	if out, err = cli.Exec(context.Background(), in); err != nil {
		panic(err)
	}
	if err = plugin.UnmarshalExecRsp(out, rsp); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", rsp)
}
