package RateLimiter

import (
	"math"
	"time"
)

type TokenBucket struct {
	tokens         float64
	maxTokens      float64
	refillRate     float64
	lastRefillTime time.Time
}

func NewTokenBucket(maxTokens float64, refillRate float64) *TokenBucket {
	return &TokenBucket{
		tokens:         maxTokens,
		maxTokens:      maxTokens,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

func (t *TokenBucket) refill() {
	currentTime := time.Now()
	duration := currentTime.Sub(t.lastRefillTime)
	tokensToAdd := t.refillRate * duration.Seconds()
	t.tokens = math.Min(t.tokens+tokensToAdd, t.maxTokens)
	t.lastRefillTime = currentTime
}

func (t *TokenBucket) Request(tokens float64) bool {
	t.refill()
	if tokens < t.tokens {
		t.tokens -= tokens
		return true
	}
	return false
}
