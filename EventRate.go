package eventBuckets

// EventRate - calculates the current failure rate based on the bucket size and time window
func (fb *Bucket) EventRate() float64 {
	fb.lock.Lock()
	defer fb.lock.Unlock()

	count := len(fb.events)
	rate := float64(count) / fb.timeWindow.Seconds()
	return rate
}
