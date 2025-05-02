package main

import "testing"

func TestSearch(t *testing.T) {
	key := "test"
	value := "this is just a test"

	dictionary := Dictionary{key: value}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(key)
		expected := value

		assertStrings(t, expected, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		expected := ErrNotFound

		assertError(t, expected, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"

		dictionary := Dictionary{}
		dictionary.Add(key, value)

		assertDefinition(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"

		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new test")

		assertError(t, ErrWordExists, err)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"

		dictionary := Dictionary{key: value}
		newDefinition := "new definition"

		err := dictionary.Update(key, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"

		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		dictionary := Dictionary{key: "test definition"}

		dictionary.Delete(key)

		_, err := dictionary.Search(key)
		assertError(t, err, ErrNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, expected, got string) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func assertError(t testing.TB, expected, got error) {
	t.Helper()

	if expected != got {
		t.Errorf("expected error %q, got %q", expected, got)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, expected string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, expected, got)
}
