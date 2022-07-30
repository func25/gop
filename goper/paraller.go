package goper

import (
	"sync"
	"time"

	"github.com/func25/mafu"
)

type paraller struct {
	batches []func()

	w        sync.WaitGroup
	elapsed  time.Duration
	batchNum int
}

func Paraller(batch int) *paraller {
	return &paraller{batchNum: batch}
}

func (p *paraller) AddWorks(funcs ...func()) {
	p.batches = append(p.batches, funcs...)
}

func (p *paraller) ClearWork() {
	p.batches = make([]func(), 0)
}

func (p *paraller) Execute() {
	start := time.Now()

	for i := 0; i < len(p.batches); {
		des := len(p.batches)
		if p.batchNum > 0 {
			des = mafu.Min(des, i+p.batchNum)
		}

		for ; i < des; i++ {
			go p.executeFunc(i)
		}

		p.w.Wait()
	}

	p.elapsed = time.Since(start)
}

func (p *paraller) executeFunc(i int) {
	p.w.Add(1)
	defer p.w.Done()
	p.batches[i]()
}

func (p *paraller) Report() time.Duration {
	return p.elapsed
}
