package selectfun

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) (winner string) {
	duration1 := measureResponseTime(url1)
	duration2 := measureResponseTime(url2)

	if duration1 < duration2 {
		return url1
	}

	return url2
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func RacerConcurrent(url1, url2 string) (winner string) {
	// select permite escuchar en múltiples canales a la vez y ejecutar el que primero responda
	select {
	case <-ping(url1):
		return url1
	case <-ping(url2):
		return url2
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
