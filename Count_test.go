package eventBuckets

import (
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	bucket := NewBucket(10 * time.Second)
	if count := bucket.Count(); count != 0 && count != bucket.totalCount {
		t.Fatal("count 0 expected")
	}
	bucket.AddEvent()
	bucket.AddEvent()
	bucket.AddEvent()
	if count := bucket.Count(); count != 3 && count != bucket.totalCount {
		t.Fatal("count 3 expected")
	}
}
