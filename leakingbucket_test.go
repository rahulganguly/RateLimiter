package RateLimiter

import (
	"testing"
	"time"
)

func TestShouldPassAllRequestsAsNumberOfTokensRequestedIsEqualToCapacity(t *testing.T) {
	var leakingBucket = NewLeakingBucket(5, time.Second)

	for i := 0; i < 5; i++ {
		isAcquired := leakingBucket.Acquire()
		if !isAcquired {
			t.Fail()
		}
	}

}

func TestShouldFailLastRequestsAsNumberOfTokensRequestedIsGreaterThanCapacity(t *testing.T) {
	var leakingBucket = NewLeakingBucket(5, time.Second)

	for i := 0; i < 5; i++ {
		isAcquired := leakingBucket.Acquire()
		if !isAcquired {
			t.Fail()
		}
	}

	if leakingBucket.Acquire() {
		t.Fail()
	}
}
