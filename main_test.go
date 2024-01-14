package main

import "testing"

func TestSpeak(t *testing.T) {
	expected := "Hello World!"
	result := Speak()
	if result != expected {
		t.Errorf("unexpected result '%s' expected: '%s'", result, expected)
	}
}
