package eventBuckets

import (
	"testing"
	"time"
)

func TestNewBucket(t *testing.T) {
	bucket := NewBucket(10 * time.Second)
	if len(bucket.events) != 0 {
		t.Fatal("zero events expected")
	}
	if bucket.timeWindow != 10*time.Second {
		t.Fatal("timeWindow expected")
	}
}
