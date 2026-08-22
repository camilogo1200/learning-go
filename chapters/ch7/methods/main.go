package main

import "fmt"

type Calculator struct {
	total  float32
	memory []string
}

// Methods are functions too
func (c *Calculator) Set(v float32) {
	c.total = v
}

func (c *Calculator) Value() float32 {
	return c.total
}

func (c *Calculator) Add(v float32) float32 {
	result := c.Value() + v
	c.Set(result)
	return c.Value()
}

func main() {

	calculator := Calculator{}
	calculator.Set(10)

	//you can assign a method to a parameter
	f2 := calculator.Add

	fmt.Printf("current value  = {%f}\n", calculator.Value())
	fmt.Printf("current Value = {%f}", f2(10))
}
