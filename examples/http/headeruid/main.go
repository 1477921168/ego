package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"

	"github.com/1477921168/ego"
	"github.com/1477921168/ego/core/elog"
	"github.com/1477921168/ego/core/etrace"
	"github.com/1477921168/ego/core/transport"
	"github.com/1477921168/ego/server/egin"
)

// export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	if err := ego.New().Serve(func() *egin.Component {
		server := egin.Load("server.http").Build()

		server.GET("/panic", func(ctx *gin.Context) {
			<-ctx.Request.Context().Done()
			panic(ctx.Request.Context().Err())
		})

		server.GET("/200", func(ctx *gin.Context) {
			<-ctx.Request.Context().Done()
			fmt.Println(ctx.Request.Context().Err())
			ctx.String(200, "hello")
		})

		server.GET("/hello", func(ctx *gin.Context) {
			pCtx := transport.WithValue(ctx.Request.Context(), "X-Ego-Uid", 9527)
			ctx.Request = ctx.Request.WithContext(pCtx)
			// Get traceId from Request's context
			// span, _ := etrace.StartSpanFromContext(ctx.Request.Context(), "Handle: /Hello")
			// defer span.Finish()

			_, span := etrace.NewTracer(trace.SpanKindServer).Start(ctx.Request.Context(), "Handle: /Hello", nil)
			defer span.End()

			ctx.JSON(200, "Hello client: "+ctx.GetHeader("app"))
		})

		return server
	}()).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}
