package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Rolandas")
	want := "Hello, Rolandas"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
