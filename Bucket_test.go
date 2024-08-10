package eventBucket

import (
	"testing"
	"time"
)

func TestBucket_struct(t *testing.T) {
	n := Bucket{
		events:     []time.Time{time.Now()},
		timeWindow: 1 * time.Second,
	}
	n.lock.Lock()
	n.lock.Unlock()
}
