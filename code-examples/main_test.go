package main

import "testing"

func TestHelloGo(t *testing.T) {
	expected := "Hello, Go!"
	if result := HelloGo(); result != expected {
		t.Errorf("HelloGo() = %q; want %q", result, expected)
	}
}
