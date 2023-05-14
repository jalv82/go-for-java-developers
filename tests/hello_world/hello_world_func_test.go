package main

import "testing"

func TestMy1stFunction(t *testing.T) {
	// Given
	word := "Hello"

	// When
	got := my1stFunction(word)

	// Then
	want := "Hello world!"
	if got != want {
		t.Fatalf("Error. I want `%s` but I got `%s`", want, got)
	}
}
