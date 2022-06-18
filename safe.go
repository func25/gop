package gop

import (
	"fmt"
	"runtime/debug"
)

func SafeGo(f func()) {
	defer Recover()
	f()
}

func Recover() {
	if r := recover(); r != nil {
		fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
	}
}
