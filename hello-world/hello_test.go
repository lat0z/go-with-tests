package main 

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage :=func(t testing.TB, got, want string){
		//t.Helper is needed to tell the test suite this is a helper.
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q ", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T){
		assertCorrectMessage(t, Hello("me",""), "Hello, me")
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got :=Hello("","")
		want := "Hello, world"
		if got !=want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("In spanish ", func(t *testing.T){
		got :=Hello("me","Spanish")
		want:= "Hola, me"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In french ", func(t *testing.T){
		got :=Hello("me","French")
		want:= "Bonjour, me"
		assertCorrectMessage(t, got, want)
	})


}