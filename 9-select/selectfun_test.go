package selectfun

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("race without concurrency", func(t *testing.T) {
		slowServer := makeDelayedServer(200 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expected := fastURL
		got := Racer(slowURL, fastURL)

		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})

	t.Run("race with concurrency", func(t *testing.T) {
		slowServer := makeDelayedServer(200 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expected := fastURL
		got := RacerConcurrent(slowURL, fastURL)

		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
