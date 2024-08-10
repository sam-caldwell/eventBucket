package eventBucket

import (
	"time"
)

// AddEvent - records a failure event at the current time
func (fb *Bucket) AddEvent() {
	fb.lock.Lock()
	fb.events = append(fb.events, time.Now())
	fb.totalCount++
	fb.lock.Unlock()
}
