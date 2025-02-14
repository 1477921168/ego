package main

import (
	"errors"

	"github.com/1477921168/ego"
	"github.com/1477921168/ego/core/elog"
)

func main() {
	err := ego.New().Invoker(func() error {
		elog.Info("logger info", elog.String("gopher", "ego"), elog.String("type", "command"))
		return errors.New("i am panic")
	}).Run()
	if err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}
