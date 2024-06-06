package RateLimiter

import (
	"testing"
	"time"
)

func TestNoRequestsFailIfWeAreConsumingOneRequestPerSecond(t *testing.T) {
	var tokenBucket = NewTokenBucket(10, 1)
	isError := false
	for i := 0; i < 10; i++ {
		if !tokenBucket.Request(1) {
			isError = true
		}
		time.Sleep(1000 * time.Millisecond)
	}
	if isError {
		t.Fail()
	}
}

func TestRequestsFailIfWeAreConsumingRequestFasterThanRefillRate(t *testing.T) {
	var tokenBucket = NewTokenBucket(20, 1)
	for i := 0; i < 20; i++ {
		if !tokenBucket.Request(10) {

		}
		time.Sleep(10 * time.Millisecond)
	}
}
