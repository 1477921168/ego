package egin

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/1477921168/ego/examples/helloworld"
)

type GreeterMock struct{}

func (mock GreeterMock) SayHello(context context.Context, request *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	return &helloworld.HelloResponse{
		Message: "hello",
	}, nil
}

func TestGRPCProxyWrapper(t *testing.T) {
	router := gin.New()
	mock := GreeterMock{}
	router.POST("/", GRPCProxy(mock.SayHello))

	// RUN
	w := performRequest(router, "POST", "/")

	assert.Equal(t, 200, w.Code)
}
