package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Gennady", "English")
		want := "Hello, Gennady"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' if an emptry string is provided", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Louis", "French")
		want := "Bonjour, Louis"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello in Turkish", func(t *testing.T) {
		got := Hello("Esra", "Turkish")
		want := "Merhaba, Esra"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
