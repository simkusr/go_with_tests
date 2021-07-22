package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Rolandas", "")
		want := "Hello, Rolandas"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say `Hello, world` when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Rolai", "French")
		want := "Bonjoure, Rolai"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Lithuanian", func(t *testing.T) {
		got := Hello("Rolandai", "Lithuanian")
		want := "Labas, Rolandai"
		assertCorrectMessage(t, got, want)
	})
}
