package test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"

	cegrpc "github.com/1477921168/ego/client/egrpc"
	"github.com/1477921168/ego/core/eerrors"
	"github.com/1477921168/ego/internal/test/errcode"
	"github.com/1477921168/ego/internal/test/helloworld"
	"github.com/1477921168/ego/server/egrpc"
)

var svc *egrpc.Component

func init() {
	svc = egrpc.DefaultContainer().Build(egrpc.WithNetwork("bufnet"))
	helloworld.RegisterGreeterServer(svc, &Greeter{})
	err := svc.Init()
	if err != nil {
		log.Fatalf("init exited with error: %v", err)
	}
	go func() {
		err = svc.Start()
		if err != nil {
			log.Fatalf("init start with error: %v", err)
		}
	}()
}

func TestGrpcError(t *testing.T) {
	resourceClient := cegrpc.DefaultContainer().Build(cegrpc.WithBufnetServerListener(svc.Listener().(*bufconn.Listener)))
	ctx := context.Background()
	client := helloworld.NewGreeterClient(resourceClient.ClientConn)
	_, err := client.SayHello(ctx, &helloworld.HelloRequest{})
	egoErr := eerrors.FromError(err)
	assert.ErrorIs(t, egoErr, errcode.ErrInvalidArgument())
	assert.Equal(t, "name is empty", egoErr.GetMessage())
}

func TestGrpcOk(t *testing.T) {
	resourceClient := cegrpc.DefaultContainer().Build(cegrpc.WithBufnetServerListener(svc.Listener().(*bufconn.Listener)))
	ctx := context.Background()
	client := helloworld.NewGreeterClient(resourceClient.ClientConn)
	resp, err := client.SayHello(ctx, &helloworld.HelloRequest{
		Name: "Ego",
	})
	assert.NoError(t, err)
	assert.True(t, proto.Equal(&helloworld.HelloResponse{
		Message: "Hello Ego",
	}, resp))

}

// Greeter ...
type Greeter struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello ...
func (g Greeter) SayHello(context context.Context, request *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	if request.Name == "" {
		return nil, errcode.ErrInvalidArgument().WithMessage("name is empty")
	}

	return &helloworld.HelloResponse{
		Message: "Hello " + request.Name,
	}, nil
}
