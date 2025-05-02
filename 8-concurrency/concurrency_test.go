package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) result {
	time.Sleep(20 * time.Millisecond)
	return result{"", true}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func mockWebsiteChecker(url string) result {
	return result{url, url != "waat://furhurterwe.geds"}
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	expect := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("expected %v, got %v", expect, got)
	}
}
