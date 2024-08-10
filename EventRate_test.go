package eventBuckets

import (
	"testing"
	"time"
)

func TestBucket_EventRate(t *testing.T) {
	// Create a new bucket with a time window of 10 seconds
	timeWindow := 1 * time.Second
	bucket := NewBucket(timeWindow) // Replace with your actual constructor

	// Add 5 events spaced by 2 seconds each
	bucket.AddEvent() // Event 1
	bucket.AddEvent() // Event 1
	bucket.AddEvent() // Event 1
	time.Sleep(200 * time.Microsecond)
	bucket.AddEvent() // Event 2
	time.Sleep(200 * time.Microsecond)
	bucket.AddEvent() // Event 3
	time.Sleep(200 * time.Microsecond)
	bucket.AddEvent() // Event 4
	time.Sleep(200 * time.Microsecond)
	bucket.AddEvent() // Event 5
	time.Sleep(200 * time.Microsecond)

	// Calculate event rate
	rate := bucket.EventRate()

	expectedRate := float64(7.0) / bucket.timeWindow.Seconds() // 5 events in 10 seconds

	if rate != expectedRate {
		t.Fatalf("Rate error\n"+
			"expected rate: %f\n"+
			"  actual rate: %f", expectedRate, rate)
	}

	// Wait for time to exceed the time window and check the rate again
	time.Sleep(10 * time.Second) // Wait long enough to prune old events
	rate = bucket.EventRate()

	expectedRate = 0.0
	if rate != expectedRate {
		t.Fatalf("Final Rate error\n"+
			"expected rate %f\n"+
			"got %f (%d events)", expectedRate, rate, len(bucket.events))
	}
}
