package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mick")
	want := "Hello, Mick"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
