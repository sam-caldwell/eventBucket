package eventBucket

import (
	"time"
)

// pruneOldEvent - removes failures that are older than the time window
//
// Note this is not thread-safe. It does not lock the bucket
// and assumes the caller has done any locking, etc.
func (fb *Bucket) pruneOldEvent() {
	fb.lock.Lock()
	defer fb.lock.Unlock()

	currentTime := time.Now()
	threshold := currentTime.Add(-fb.timeWindow)
	for _, failureTime := range fb.events {
		if failureTime.After(threshold) {
			if len(fb.events) > 1 {
				fb.events = append([]time.Time{}, fb.events[1:]...)
			} else {
				fb.events = []time.Time{}
			}
		}
	}
}
