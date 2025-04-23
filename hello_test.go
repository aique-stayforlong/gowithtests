package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Aique")
	want := "Hello, Aique"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}