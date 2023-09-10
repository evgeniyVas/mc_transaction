package worker

import (
	"context"
	"github.com/mc_transaction/internal/logger"
	"sync"
	"time"
)

type Workers struct {
	name    string
	ctx     context.Context
	cancelF context.CancelFunc
	wg      *sync.WaitGroup

	cnt int

	ticker *time.Ticker

	exec  executor
	again chan struct{}
}

type executor func(ctx context.Context)

func New(name string, count int, period time.Duration, work executor) *Workers {
	return &Workers{
		name:   name,
		wg:     &sync.WaitGroup{},
		cnt:    count,
		exec:   work,
		again:  make(chan struct{}, 1),
		ticker: time.NewTicker(period),
	}
}

func (s *Workers) Start(ctx context.Context) {
	s.ctx, s.cancelF = context.WithCancel(ctx)
	s.wg.Add(s.cnt)
	for i := 0; i < s.cnt; i++ {
		go s.start(s.ticker)
	}
}

func (s *Workers) Close() {
	logger.Info("%s is closing " + s.name)
	s.cancelF()
	s.wg.Wait()
	s.ticker.Stop()
	close(s.again)
}

func (s *Workers) start(ticker *time.Ticker) {
	defer s.wg.Done()
	s.Trigger()
	for {
		select {
		case <-s.ctx.Done():
			return
		case <-s.again:
			logger.Debug("%s is trying again " + s.name)
			select {
			case <-s.ctx.Done():
				return
			default:
			}
			logger.Debug("%s is working" + s.name)
			s.exec(s.ctx)
			continue
		case <-ticker.C:
			logger.Debug("%s 's ticker " + s.name)
			s.Trigger()
		}
	}

}

func (s *Workers) Trigger() {
	select {
	case s.again <- struct{}{}:
	default:
	}
}

func (s *Workers) FuncNext() func() {
	return s.Trigger
}
