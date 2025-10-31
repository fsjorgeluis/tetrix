package infrastructure

import "time"

type Timer struct {
	ticker *time.Ticker
	stop   chan struct{}
}

func NewTimer(interval time.Duration) *Timer {
	return &Timer{
		ticker: time.NewTicker(interval),
		stop:   make(chan struct{}),
	}
}

func (t *Timer) Start(tickChan chan struct{}) {
	go func() {
		for {
			select {
			case <-t.ticker.C:
				tickChan <- struct{}{}
			case <-t.stop:
				t.ticker.Stop()
				return
			}
		}
	}()
}

func (t *Timer) Stop() {
	close(t.stop)
}
