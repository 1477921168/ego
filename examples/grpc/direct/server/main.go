package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/1477921168/ego"
	"github.com/1477921168/ego/core/elog"
	"github.com/1477921168/ego/examples/helloworld"
	"github.com/1477921168/ego/server"
	"github.com/1477921168/ego/server/egovernor"
	"github.com/1477921168/ego/server/egrpc"
)

// export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	if err := ego.New().Serve(func() server.Server {
		component := egrpc.Load("server.grpc").Build()
		helloworld.RegisterGreeterServer(component.Server, &Greeter{server: component})
		return component
	}(), egovernor.Load("server.governor").Build()).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}

// Greeter ...
type Greeter struct {
	server *egrpc.Component
	helloworld.UnimplementedGreeterServer
}

// SayHello ...
func (g Greeter) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	if request.Name == "error" {
		return nil, status.Error(codes.Unavailable, "error")
	}
	egrpc.CtxStoreSet(ctx, "x-biz-guid", "123")
	egrpc.CtxStoreSet(ctx, "x-biz-uid", "100")
	// header := metadata.Pairs("x-header-key", "val")
	// err := grpc.SendHeader(context, header)
	// if err != nil {
	//	return nil, fmt.Errorf("set header fail, %w", err)
	// }
	// go func() {
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			fmt.Println(ctx.Err())
	//			return
	//		}
	//	}
	// }()

	<-ctx.Done()
	return nil, ctx.Err()

	// time.Sleep(xtime.Duration("2s"))
	return &helloworld.HelloResponse{
		Message: "Hello EGO, I'm " + g.server.Address(),
	}, nil
}
