package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people ", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})
	t.Run("say Hello, world when an empty string is passed", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

}

func TestTemp(t *testing.T) {
	got := Hello("world", "French")
	want := "Hello world"
	assertCorrectMessage(t, got, want)

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
