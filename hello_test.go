package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Aique")
	want := "Hello, Aique"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHelloDefault(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Aique")
		want := "Hello, Aique"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // permite hacer referencia a la línea del test si algo falla

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}