package main

import (
	"time"

	"github.com/1477921168/ego"
	"github.com/1477921168/ego/core/econf"
	"github.com/1477921168/ego/core/elog"
)

// export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	if err := ego.New(ego.WithHang(true)).Invoker(func() error {
		go func() {
			// 循环打印配置
			for {
				time.Sleep(3 * time.Second)
				peopleName := econf.GetString("people.name")
				elog.Info("people info", elog.String("name", peopleName), elog.String("type", "onelineByFileWatch"))
			}
		}()
		return nil
	}).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}
