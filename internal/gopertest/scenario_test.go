package goptest

import (
	"fmt"
	"testing"
	"time"
)

type fakeInt int

func (f fakeInt) println1() {
	fmt.Println(f)
}

func (f *fakeInt) println2() {
	fmt.Println(*f)
}

func Test1(t *testing.T) {
	i := 0
	defer func() {
		fmt.Println(i)
	}()
	i++
}

func Test2(t *testing.T) {
	i := 0
	defer func(i int) {
		fmt.Println(i)
	}(i)
	i++
}

func Test3(t *testing.T) {
	var i fakeInt = 0
	defer i.println1()
	i++
}

func Test4(t *testing.T) {
	var i = 1
	go func() {
		fmt.Println("from goroutine:", i)
	}()
	i++
	time.Sleep(time.Second)
}

func Test5(t *testing.T) {
	var i = 1
	go func(i int) {
		fmt.Println("from goroutine:", i)
	}(i)
	i++
	time.Sleep(time.Second)
}

func Test6(t *testing.T) {
	var i = 1
	go func() {
		fmt.Println("start 1")
		for {
			i++
		}
	}()

	time.Sleep(2000 * time.Millisecond)

	go func() {
		fmt.Println("from goroutine:", &i)
	}()

	fmt.Println("from main:", i)

	time.Sleep(time.Second)
}

func Test6_1(t *testing.T) {
	var x = 1
	var i = &x
	go func() {
		fmt.Println("start 1")
		for {
			*i++
		}
	}()

	time.Sleep(2000 * time.Millisecond)

	go func() {
		fmt.Println("from goroutine:", x)
	}()

	fmt.Println("from main:", x)

	time.Sleep(time.Second)
}

func Test6_2(t *testing.T) {
	var x = 1
	var i = &x
	go func() {
		fmt.Println("start 1")
		for {
			*i++
		}
	}()

	time.Sleep(2000 * time.Millisecond)

	go func() {
		*i++
		fmt.Println("from goroutine:", x)
	}()

	fmt.Println("from main:", x)

	time.Sleep(time.Second)
}

func Test8(t *testing.T) {
	var i = 1

	defer func() {
		fmt.Println("from goroutine:", i)
	}()

	defer func() {
		i = 3
	}()

	fmt.Println("from main:", i)

	time.Sleep(time.Second)
}
