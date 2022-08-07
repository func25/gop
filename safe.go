package gop

import (
	"fmt"
	"reflect"
	"runtime/debug"
)

// SafeGo is designed for preventing panic and returning error instead
func SafeGo(f func()) (err error) {
	defer RecoverErr(&err)

	f()
	return
}

// Recover recovers from panicking and print error + its stacktrace
func Recover() {
	if r := recover(); r != nil {
		err := fmt.Errorf("panic: %v\nstacktrace from panic: %v", r, string(debug.Stack()))
		fmt.Println(err.Error())
	}
}

// RecoverErr returns error that it caught while recovering from panic process
func RecoverErr(err *error) {
	if r := recover(); r != nil {
		catch := fmt.Errorf("panic: %v\nstacktrace from panic: %v", r, string(debug.Stack()))
		reflect.ValueOf(err).Elem().Set(reflect.ValueOf(catch))
	}
}
