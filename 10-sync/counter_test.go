package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, 3, counter)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		// espera a que finalice una colección de goroutines
		var wg sync.WaitGroup

		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		// bloquea la ejecución hasta que no queden elementos
		wg.Wait()

		assertCounter(t, wantedCount, counter)
	})
}

func assertCounter(t testing.TB, expected int, c *Counter) {
	t.Helper()

	if expected != c.Value() {
		t.Errorf("expected %d, got %d", expected, c.Value())
	}
}
