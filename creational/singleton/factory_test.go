package singleton_test

import (
	"design_pattern/creational/singleton/fraud"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestFruadEngine(t *testing.T) {
	t.Run("Test single construct", func(t *testing.T) {
		var wg sync.WaitGroup
		start := time.Now()

		// Simulate 50 concurrent requests hitting the endpoint
		for id := 0; id < 50; id++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				// All these should block until the ONE init is done
				e := fraud.GetFraudEngine()
				fmt.Printf("Req %d: Engine Loaded? %v\n", id, e != nil)
			}(id)
		}

		wg.Wait()
		fmt.Printf("Total Time: %v (Should be ~2s, not 100s)\n", time.Since(start))
	})
}