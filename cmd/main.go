package main

import (
	"fmt"
	"time"

	"github.com/AnhellO/ticker/internal/file"
	"github.com/AnhellO/ticker/internal/ticker"
)

// cache represents the in-memory cache used for r+w
var cache map[string]interface{}

func main() {
	// cache initialization
	cache = map[string]interface{}{
		"file": file.Users{},
	}

	refresher := ticker.NewTickerRefresher(6*time.Second, "wrong-key")
	go refresher.Refresh(cache)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%+v\n", cache)
			time.Sleep(10 * time.Second)
		}
	}()

	time.Sleep(100 * time.Second)
}
