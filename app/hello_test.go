package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("with name", func(t *testing.T) {
		got := Hello("edgar")
		want := "Hello, edgar"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("with empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
