package pokeapi

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/Relevantfender/pokedexcli/internal/pokecache"
)

func TestConcurrentAccess(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Second)
	var wg sync.WaitGroup

	// Add a bunch of items concurrently
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i)
			val := []byte(fmt.Sprintf("val-%d", i))
			cache.Add(key, val)
		}(i) // Pass i as an argument to the goroutine
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Verify all items were added correctly
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key-%d", i)
		val, ok := cache.Get(key)
		if !ok {
			t.Errorf("Expected to find key %s", key)
			continue
		}
		expected := fmt.Sprintf("val-%d", i)
		if string(val) != expected {
			t.Errorf("Expected value %s for key %s, got %s", expected, key, string(val))
		}
	}
}
