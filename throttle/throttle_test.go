//
// throttle
// Copyright © 2023 Space Code. All rights reserved.
//

package throttle_test

import (
	"testing"
	"time"

	"github.com/space-code/throttle/throttle"
)

const throttleDuration = time.Second * 1

func TestThrottle(t *testing.T) {
	throttler := throttle.New(throttleDuration)

	x := 0
	f := func() { x++ }

	for i := 0; i < 3; i++ {
		throttler.Do(f)
	}

	time.Sleep(throttleDuration)

	if x != 1 {
		t.Errorf("Expected x to be 1, but got %d", x)
	}
}
