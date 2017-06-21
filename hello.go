package hello

import (
	"fmt"
)

// Greeter provides invidualized greeting messages
type Greeter interface {

	// Greet returns an individual greeting message
	Greet(name string) string
}

// greeter is the default implementation of Greeter
type greeter struct {
	greetingTemplate string
}

// Returns an instance to the  default Greeter
func DefaultGreeter() Greeter {
	return GreeterWith("")
}

// Returns a Greeter instance with the provided greeting template
func GreeterWith(greetingTemplate string) Greeter {

	if greetingTemplate != "" {
		return &greeter{greetingTemplate}
	}

	return &greeter{"Hello %s"}
}

// Greet returns a simple individualized greeting
func (g *greeter) Greet(name string) string {

	if name == "" {
		name = "World"
	}

	return fmt.Sprintf(g.greetingTemplate, name)
}
