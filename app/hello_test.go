package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("edgar")
	want := "Hello, edgar"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
