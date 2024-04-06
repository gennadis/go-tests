package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Gennady")
	want := "Hello, Gennady"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
