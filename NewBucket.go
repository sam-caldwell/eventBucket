package eventBucket

import (
	"time"
)

// NewBucket creates a new Bucket and starts the background pruning goroutine
func NewBucket(timeWindow time.Duration) *Bucket {
	b := &Bucket{
		events:     []time.Time{},
		timeWindow: timeWindow,
		stopChan:   make(chan struct{}),
	}
	b.startPruning()
	return b
}
