package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	result, reminder, error := divAndRemainder(10, 0)
	if error != nil {
		fmt.Printf("error = {%v}\n", error.Error())
	}
	fmt.Printf("result = {%f}, reminder = {%f}\n", result, reminder)

	result, reminder, error = divAndRemainder(10, 30)
	if error != nil {
		fmt.Printf("error = {%v}\n", error.Error())
	}
	fmt.Printf("result = {%f}, reminder = {%f}\n", result, reminder)

	//you can named result values check the sum function
	sum, err := sum(10, 55)

	fmt.Printf("sum = {%d}, err ={%v}\n", sum, err)

	//blank return are discouraged, however you can use them
	sum, err = sum2(10, 20)
	fmt.Printf("sum2 = {%d}, err ={%v}\n", sum, err)

}

func divAndRemainder(n float32, d float32) (float32, float32, error) {
	if d == 0 {
		return -1, -1, errors.New("can't do operation with zero denominator")
	}
	return n / d, float32(math.Mod(float64(n), float64(d))), nil
}

//named return values

func sum(a int, b int) (result int, error error) {
	return a + b, nil
}

func sum2(a int, b int) (result int, error error) {
	result = a + b
	return
}
