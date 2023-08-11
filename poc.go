package ticker

import "time"

// any is an alias for interface{} and is equivalent to interface{} in all ways.
type any = interface{}

// TickerRefresher represents a componen able to dynamically refresh its data every certain time.
type TickerRefresher interface {
	// SetData sets the data that needs to be refreshed every certain time
	SetData(data any)

	// SetTime sets the time duration needed for the ticker
	SetTime(duration time.Duration)

	// Refresh contains the necessary logic to refresh the data
	Refresh() (any, error)
}
