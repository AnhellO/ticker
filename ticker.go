// Package ticker implements a ticker that can be stopped and re-started.
package ticker

import (
	"sync"
	"time"
)

// A CustomTicker holds a channel that delivers ticks at intervals.
type CustomTicker struct {
	Ticks    chan time.Time // The channel on which ticks are delivered.
	duration time.Duration
	mu       sync.Mutex
	running  bool
	stop     chan bool
}

// New returns a new ticker that ticks every d seconds. It adjusts the
// intervals or drops ticks to make up for slow receivers. The ticker
// is initially in the stopped state.
func New(duration time.Duration) *CustomTicker {
	if duration <= 0 {
		panic("ticker: non-positive duration")
	}
	return &CustomTicker{
		Ticks:    make(chan time.Time),
		duration: duration,
		running:  false,
		stop:     make(chan bool),
	}
}

// Start (re-)starts the ticker. Ticks will be delivered on the ticker's
// channel until Stop is called.
func (t *CustomTicker) Start() {
	t.mu.Lock()
	defer t.mu.Unlock()
	if !t.running {
		go t.loop()
		t.running = true
	}
}

// Stop stops the ticker. No ticks will be delivered on the ticker's channel
// after Stop returns and before Start is called again.
func (t *CustomTicker) Stop() {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.running {
		t.stop <- true
		t.running = false
	}
}

// Stopped returns whether the ticker is stopped.
func (t *CustomTicker) Stopped() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return !t.running
}

func (t *CustomTicker) loop() {
	tk := time.NewTicker(t.duration)
	for {
		select {
		case tm := <-tk.C:
			select {
			case t.Ticks <- tm:
			default:
			}
		case <-t.stop:
			tk.Stop()
			return
		}
	}
}
