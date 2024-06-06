package RateLimiter

import "time"

type LeakingBucket struct {
	capacity   int
	rate       time.Duration
	tokens     int
	stopTicker chan bool
}

func NewLeakingBucket(capacity int, rate time.Duration) *LeakingBucket {
	bucket := &LeakingBucket{
		capacity:   capacity,
		rate:       rate,
		tokens:     capacity,
		stopTicker: make(chan bool),
	}
	go bucket.startLeaking()
	return bucket
}

func (bucket *LeakingBucket) startLeaking() {
	ticker := time.NewTicker(bucket.rate)
	for {
		select {
		case <-ticker.C:
			if bucket.tokens < bucket.capacity {
				bucket.tokens++
			}
		case <-bucket.stopTicker:
			ticker.Stop()
			return
		}
	}
}

func (bucket *LeakingBucket) Stop() {
	bucket.stopTicker <- true
}
func (bucket *LeakingBucket) Acquire() bool {
	if bucket.tokens > 0 {
		bucket.tokens--
		return true
	}
	return false
}
