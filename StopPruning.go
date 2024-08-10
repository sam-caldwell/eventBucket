package eventBuckets

// StopPruning stops the background pruning goroutine
func (fb *Bucket) StopPruning() {
	close(fb.stopChan)
	fb.ticker.Stop()
}
