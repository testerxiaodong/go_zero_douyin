package main

import (
	"flag"
	"fmt"
	"go-zero-douyin/common/interceptor/rpcServer"

	"go-zero-douyin/apps/social/cmd/rpc/internal/config"
	"go-zero-douyin/apps/social/cmd/rpc/internal/server"
	"go-zero-douyin/apps/social/cmd/rpc/internal/svc"
	"go-zero-douyin/apps/social/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/social.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterSocialServer(grpcServer, server.NewSocialServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//rpc logger
	s.AddUnaryInterceptors(rpcServer.ErrTransLogInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
