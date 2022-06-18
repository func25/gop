package gopbot

import (
	"time"
)

type Worker[T any] struct {
	Funcs []FuncDef[T]
	state T
}

type FuncDef[T any] struct {
	F                 func(T) (T, error)
	NextConditionFunc func(T) bool
	TimeoutFunc       func(T, time.Duration) bool
}

func (w Worker[T]) Do(inp T) (T, error) {
	var err error
	w.state = inp
	for i := range w.Funcs {
		w.state, err = w.Funcs[i].F(w.state)
		if err != nil {
			return w.state, err
		}

		if w.Funcs[i].NextConditionFunc != nil {
			startTime := time.Now()
			for !w.Funcs[i].NextConditionFunc(w.state) {
				waitTime := time.Now()
				if w.Funcs[i].TimeoutFunc != nil && w.Funcs[i].TimeoutFunc(w.state, waitTime.Sub(startTime)) {
					break
				}
			}
		}
	}

	return w.state, err
}

// functask: x2, x3, x4 speed
