//
// throttle
// Copyright © 2023 Space Code. All rights reserved.
//

package throttle

import (
	"sync"
	"time"
)

// Throttler is an interface for rate limiting.
type Throttler interface {
	Do(f func())
}

// throttler is a structure for task throttling.
type throttler struct {
	interval time.Duration
	once     sync.Once
	m        sync.Mutex
}

// New creates a new instance of Throttler.
func New(interval time.Duration) Throttler {
	return &throttler{
		interval: interval,
	}
}

// Do executes a task through throttling, ensuring that the task is only executed once
// within a specified time interval.
func (th *throttler) Do(f func()) {
	th.m.Lock()
	defer th.m.Unlock()

	th.once.Do(func() {
		go func() {
			th.m.Lock()
			defer th.m.Unlock()

			time.Sleep(th.interval)
			th.once = sync.Once{}
		}()

		f()
	})
}
