package goper

import (
	"sync"
	"time"

	"github.com/func25/mafu"
)

type paraller struct {
	// batch state
	batches  []func() error
	w        sync.WaitGroup
	batchNum int

	// execution report
	elapsed time.Duration

	// error state, report
	errors      []funcErr
	errChan     chan *funcErr
	stopWhenErr bool
}

func Paraller(batch int) *paraller {
	return &paraller{batchNum: batch, errChan: make(chan *funcErr)}
}

func (p *paraller) AddWorks(funcs ...func() error) *paraller {
	p.batches = append(p.batches, funcs...)

	return p
}

func (p *paraller) ClearWork() *paraller {
	p.batches = make([]func() error, 0)

	return p
}

func (p *paraller) StopWhenError(stop bool) *paraller {
	p.stopWhenErr = stop

	return p
}

func (p *paraller) Execute() []funcErr {
	p.startReport()
	defer p.stopReport()

	for i := 0; i < len(p.batches); {
		des := len(p.batches)
		if p.batchNum > 0 {
			des = mafu.Min(des, i+p.batchNum)
		}

		for ; i < des; i++ {
			go p.executeFunc(i)
		}

		p.w.Wait()

		if p.stopWhenErr && len(p.errors) > 0 {
			return p.errors
		}
	}

	return p.errors
}

func (p *paraller) startReport() {
	start := time.Now()
	p.errors = nil

	for {
		select {
		case err := <-p.errChan:
			if err == nil {
				p.elapsed = time.Since(start)
				return
			}
			p.errors = append(p.errors, *err)
		}
	}
}

func (p *paraller) stopReport() {
	p.errChan <- nil
}

func (p *paraller) executeFunc(i int) error {
	p.w.Add(1)
	defer p.w.Done()

	if err := p.batches[i](); err != nil {
		p.errChan <- &funcErr{
			Func: p.batches[i],
			Err:  err,
		}
	}

	return nil
}

func (p *paraller) Report() time.Duration {
	return p.elapsed
}
