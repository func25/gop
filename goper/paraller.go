package goper

import (
	"sync"
	"time"
)

type paraller struct {
	batches []func()

	w       sync.WaitGroup
	elapsed time.Duration
}

func Paraller() *paraller {
	return &paraller{}
}

func (p *paraller) AddWork(funcs ...func()) {
	for i := range funcs {
		p.batches = append(p.batches, funcs[i])
	}
}

func (p *paraller) ClearWork() {
	p.batches = make([]func(), 0)
}

func (p *paraller) Execute() {
	start := time.Now()

	for i := range p.batches {
		p.execute(i)
	}

	p.w.Wait()
	p.elapsed = time.Since(start)
}

func (p *paraller) execute(i int) {
	p.w.Add(1)
	defer p.w.Done()
}

func (p *paraller) Report() time.Duration {
	return p.elapsed
}
