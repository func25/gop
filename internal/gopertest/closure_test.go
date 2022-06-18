package goptest

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestClosure_1(t *testing.T) {
	x := incrementor()
	fmt.Println(unsafe.Sizeof(x))
	for i := 0; i < 10; i++ {
		fmt.Println(x())
	}
}

func incrementor() func() int {
	i := 0
	x := 2
	return func() int {
		i++
		x++
		return i
	}
}

// func incrementor_pointer() func() int {
// 	x := 0
// 	i := &x
// 	return func() int {
// 		*i++
// 		return *i
// 	}
// }
