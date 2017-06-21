package hello

import (
	"fmt"
	"sync"
)

// Greeter provides invidualized greeting messages
type Greeter interface {

	// Greet returns an individual greeting message
	Greet(name string) string
}

// greeter is the default implementation of Greeter
type greeter struct {
}

var (
	instance Greeter
	once     sync.Once
)

// Returns the default Greeter instance
func DefaultGreeter() Greeter {

	once.Do(func() {
		instance = &greeter{}
	})

	return instance
}

// Greet returns a simple individualized greeting
func (g *greeter) Greet(name string) string {

	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("Hello %s", name)
}
