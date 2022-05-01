package goper

import "time"

type Worker[T WorkerState] struct {
	Funcs []FuncDef[T]
	state T
}

type FuncDef[T WorkerState] struct {
	F                 func(T) T
	NextConditionFunc func(T) bool
	TimeoutFunc       func(T, time.Duration) bool
}

type WorkerState interface {
	Err() error
}

func (w Worker[T]) Do(inp T) (T, error) {
	w.state = inp
	for i := range w.Funcs {
		w.state = w.Funcs[i].F(inp)
		if w.state.Err() != nil {
			break
		}

		if w.Funcs[i].NextConditionFunc != nil {
			startTime := time.Now()
			for w.Funcs[i].NextConditionFunc(w.state) {
				waitTime := time.Now()
				if w.Funcs[i].TimeoutFunc != nil && w.Funcs[i].TimeoutFunc(w.state, waitTime.Sub(startTime)) {
					break
				}
			}
		}
	}

	return w.state, w.state.Err()
}
