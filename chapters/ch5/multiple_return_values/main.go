package main

import (
	"errors"
	"fmt"
)

func main() {

	result, reminder, error := div_and_remainder(10, 0)
	if error != nil {
		fmt.Printf(" error = {%v}\n", error.Error())
		return
	}
	fmt.Printf("result = {%f}, reminder = {%f}", result, reminder)

	//you can named result values check the sum function
	sum, err := sum(10, 55)
}

func div_and_remainder(n float32, d float32) (float32, float32, error) {
	if d == 0 {
		return -1, -1, errors.New("can't do operation with zero denominator")
	}
	return 0, 0, nil
}

//named return values

func sum(a int, b int) (result int, error error) {

	return a + b, nil
}
