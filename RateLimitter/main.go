package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Visitor struct {
	LastRequest time.Time
}

type RateLimiter struct {
	visitor  map[string]Visitor
	duration time.Duration
	mu       sync.Mutex
}

func NewRateLimiter(ctx context.Context) *RateLimiter {
	rl := &RateLimiter{
		visitor:  make(map[string]Visitor),
		duration: time.Minute,
	}

	go rl.CleanUp(ctx)

	return rl
}

func (r *RateLimiter) Allow(userId string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()

	u, exist := r.visitor[userId]
	if !exist {
		r.visitor[userId] = Visitor{
			LastRequest: now,
		}
		return true
	}

	if now.Sub(u.LastRequest) < r.duration {
		return false
	}

	u.LastRequest = now
	r.visitor[userId] = u
	return true
}

func (r *RateLimiter) CleanUp(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			r.mu.Lock()
			for userId, visitor := range r.visitor {
				if time.Since(visitor.LastRequest) > r.duration {
					delete(r.visitor, userId)
				}
			}
			r.mu.Unlock()

		case <-ctx.Done():
			fmt.Println("Cleanup goroutine stopped:", ctx.Err())
			return
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	limiter := NewRateLimiter(ctx)

	userId := "user-123"

	fmt.Printf("Request 1: %v \n", limiter.Allow(userId))

	fmt.Printf("Request 2: %v \n", limiter.Allow(userId))

	time.Sleep(10 * time.Second)

	fmt.Printf("Request 3: %v \n", limiter.Allow(userId))

	time.Sleep(51 * time.Second)

	fmt.Printf("Request 3: %v \n", limiter.Allow(userId))
}
