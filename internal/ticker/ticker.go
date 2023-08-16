// Package ticker implements a ticker that can be stopped and re-started.
package ticker

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/AnhellO/ticker/internal/file"
)

// any is an alias for interface{} and is equivalent to interface{} in all ways.
type any = interface{}

// TickerRefresher represents a component able to dynamically refresh its data every certain time.
type TickerRefresher interface {
	// Refresh contains the necessary logic to refresh the data
	Refresh(cache map[string]interface{})
}

// NewTickerRefresher returns a new ticker that ticks every d seconds. It adjusts the
// intervals or drops ticks to make up for slow receivers. The ticker
// is initially in the stopped state.
func NewTickerRefresher(duration time.Duration, cacheKey string) TickerRefresher {
	if duration <= 0 {
		panic("ticker: non-positive duration")
	}

	return &customTicker{
		ticker:   time.NewTicker(duration),
		mutex:    sync.RWMutex{},
		cacheKey: cacheKey,
	}
}

// A customTicker holds a channel that delivers ticks at intervals.
type customTicker struct {
	ticker   *time.Ticker
	mutex    sync.RWMutex
	cacheKey string
}

func (t *customTicker) Refresh(cache map[string]interface{}) {
	for {
		select {
		case tick := <-t.ticker.C:
			// update cache
			fmt.Println("Update at", tick.UTC())
			t.mutex.Lock() // lock the cache before writing into it
			newData, err := file.GetData("./mock-data.json")
			if err != nil {
				log.Fatalf("error %+w", err)
				continue
			}
			cache[t.cacheKey] = newData
			t.mutex.Unlock() // unlock the cache before writing into it
		}
	}
}
