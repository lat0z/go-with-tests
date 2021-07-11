package main 

import "testing"

func TestHello(t *testing.T) {
	got := Hello("me")
	want := "Hello, me"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}