package main

import (
	RateLimiter "SystemDesign/Ratelimiter"
	_ "SystemDesign/Ratelimiter"
	"fmt"
	"time"
)

func main() {
	bucket := RateLimiter.NewLeakingBucket(5, time.Second)
	defer bucket.Stop()

	// Simulate requests
	for i := 0; i < 10; i++ {
		if bucket.Acquire() {
			fmt.Printf("Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: Denied\n", i+1)
		}
		time.Sleep(1 * time.Millisecond)
	}
}
