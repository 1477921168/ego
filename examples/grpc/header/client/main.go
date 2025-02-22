package main

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/1477921168/ego/core/transport"

	"github.com/1477921168/ego"
	"github.com/1477921168/ego/client/egrpc"
	"github.com/1477921168/ego/core/elog"
	"github.com/1477921168/ego/examples/helloworld"
)

func main() {
	if err := ego.New().Invoker(
		invokerGrpc,
		callGrpc,
	).Run(); err != nil {
		elog.Error("startup", elog.FieldErr(err))
	}
}

var grpcComp helloworld.GreeterClient

func invokerGrpc() error {
	grpcConn := egrpc.Load("grpc.test").Build()
	grpcComp = helloworld.NewGreeterClient(grpcConn.ClientConn)
	return nil
}

func callGrpc() error {
	req := http.Request{}
	parentContext := transport.WithValue(req.Context(), "X-Ego-Uid", 9527)
	var headers metadata.MD
	var trailers metadata.MD
	_, err := grpcComp.SayHello(parentContext, &helloworld.HelloRequest{
		Name: "i am client",
	}, grpc.Header(&headers), grpc.Trailer(&trailers))
	if err != nil {
		return err
	}

	spew.Dump(headers)
	spew.Dump(trailers)
	_, err = grpcComp.SayHello(parentContext, &helloworld.HelloRequest{
		Name: "error",
	})
	if err != nil {
		return err
	}
	return nil
}
