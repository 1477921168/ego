package main

import (
	"github.com/gin-gonic/gin"

	"github.com/1477921168/ego"
	"github.com/1477921168/ego/core/elog"
	"github.com/1477921168/ego/server/egin"
)

// export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	if err := ego.New().Serve(func() *egin.Component {
		server := egin.Load("server.http").Build()
		server.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, "Hello EGO")
			return
		})
		return server
	}()).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}
