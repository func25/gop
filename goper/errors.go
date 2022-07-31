package goper

import "github.com/func25/slicesol/slicesol"

type funcErr struct {
	Func func() error
	Err  error
}

type errReport []funcErr

func (report errReport) toFuncArray() []func() error {
	return slicesol.Map(report, func(elem funcErr) func() error {
		return elem.Func
	})
}
