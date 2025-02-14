package main

import (
	"github.com/gin-gonic/gin"

	"github.com/1477921168/ego"
	"github.com/1477921168/ego/core/elog"
	"github.com/1477921168/ego/server/egin"
	"github.com/1477921168/ego/server/egovernor"
)

func main() {
	if err := ego.New().
		Serve(
			egovernor.Load("server.governor").Build(),
			serverHTTP(),
		).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}

func serverHTTP() *egin.Component {
	server := egin.Load("server.http").Build()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello")
		return
	})
	return server
}
