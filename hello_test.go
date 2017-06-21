package hello

import (
	"testing"
)

func TestDefaultGreeting(t *testing.T) {

	cases := []struct{ name, expected string }{
		{"Tom", "Hello Tom"},
		{"Ralf", "Hello Ralf"},
		{"", "Hello World"},
	}

	greeter := DefaultGreeter()

	for _, tc := range cases {
		actual := greeter.Greet(tc.name)
		if actual != tc.expected {
			t.Errorf("with %s got %s, expected %s", tc.name, actual, tc.expected)
		}
	}
}

func TestCustomGreeting(t *testing.T) {

	cases := []struct{ name, expected string }{
		{"Tom", "Hi Tom"},
		{"Ralf", "Hi Ralf"},
		{"", "Hi World"},
	}

	greeter := GreeterWith("Hi %s")

	for _, tc := range cases {
		actual := greeter.Greet(tc.name)
		if actual != tc.expected {
			t.Errorf("with %s got %s, expected %s", tc.name, actual, tc.expected)
		}
	}
}
