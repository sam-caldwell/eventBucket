package eventBuckets

import (
	"sync"
	"time"
)

// Bucket - tracks failures within a specified time window and bucket size
//
// The bucket will contain a number of events in a given rolling time window
// from which we can calculate the event rate (e.g. failure rates, etc.).
type Bucket struct {
	lock       sync.Mutex
	events     []time.Time
	timeWindow time.Duration
	ticker     *time.Ticker
	totalCount uint64
	stopChan   chan struct{}
}
