package fraud

import (
	"fmt"
	"sync"
	"time"
)

var (
	fraudEngineOnce sync.Once
	fraudEngine     *FraudEngine
)

type FraudEngine struct {
	Rules map[string]bool
}

func GetFraudEngine() *FraudEngine {
	fraudEngineOnce.Do(func() {
		// Inside Do:
		// 1. fmt.Println("--- LOADING HEAVY DATA (2s) ---")
		// 2. time.Sleep(2 * time.Second)
		// 3. Initialize 'engine' struct
		fmt.Println("--- LOADING HEAVY DATA (2s) ---")
		time.Sleep(2*time.Second)
		fraudEngine = new(FraudEngine)
	})

	return fraudEngine
}
