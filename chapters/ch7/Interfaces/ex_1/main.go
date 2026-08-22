package main

import (
	"fmt"
	"time"
)

type Stringer interface {
	String() string
}

type Incrementer interface {
	Increment()
}
type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last Update: %v", c.total, c.lastUpdatedw)
}
func main() {
	// The only abstract type in Go.
	//Here's the definition of Stringer interface
	var myStringer Stringer
	var myIncrementer Incrementer

	pointerCounter := &Counter{}
	valueCounter := Counter{}

	myStringer = pointerCounter
	myStringer = valueCounter

	myIncrementer = pointerCounter
	myIncrementer = valueCounter
}
