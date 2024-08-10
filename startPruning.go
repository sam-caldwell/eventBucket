package eventBuckets

import (
	"time"
)

// startPruning starts a background goroutine that periodically prunes old events
func (fb *Bucket) startPruning() {

	fb.ticker = time.NewTicker(fb.timeWindow / 5)

	go func() {
		for {
			select {
			case <-fb.ticker.C:
				fb.pruneOldEvent()
			case <-fb.stopChan:
				return
			}
		}
	}()

}
