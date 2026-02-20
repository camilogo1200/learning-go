package main

import (
	"fmt"
	"math"
)

func main() {
	//Integer types -> Signed and Unsigned, from one to 8-bytes
	//Signed Integers
	var num1 int8 = math.MaxInt8
	var num2 int8 = math.MinInt8
	var num3 int16 = math.MaxInt16
	var num4 int16 = math.MinInt16
	var num5 int32 = math.MaxInt32
	var num6 int32 = math.MinInt32
	var num7 int64 = math.MaxInt64
	var num8 int64 = math.MinInt64

	fmt.Printf("Signed Integers: \n")
	fmt.Printf("Max Int8 value: [%d]\n", num1)
	fmt.Printf("Min Int8 value: [%d]\n", num2)
	fmt.Printf("Max Int16 value: [%d]\n", num3)
	fmt.Printf("Min Int16 value: [%d]\n", num4)
	fmt.Printf("Max Int32 value: [%d]\n", num5)
	fmt.Printf("Min Int32 value: [%d]\n", num6)
	fmt.Printf("Max Int64 value: [%d]\n", num7)
	fmt.Printf("Min Int64 value: [%d]\n", num8)

	//unsigned integers
	var num9 uint8 = math.MaxUint8
	var num10 uint16 = math.MaxUint16
	var num11 uint32 = math.MaxUint32
	var num12 uint64 = math.MaxUint64

	fmt.Printf("\n---------------------------------------------------------------------\n\n")
	fmt.Printf("Unsigned Integers: \n")
	fmt.Printf("Max UInt8 value: [%d]\n", num9)
	fmt.Printf("Max UInt16 value: [%d]\n", num10)
	fmt.Printf("Max UInt32 value: [%d]\n", num11)
	fmt.Printf("Max UInt32 value: [%d]\n", num12)
	//byte is an alias for uint8
	//int is architecture dependent, so can be int32 or int64
	//uint is architecture dependent, so can be uint32 or uint64

	//there are two special names for integer types ( rune and uintptr)
	//integer operators +, -, *, /, %
	//the result of an integer operation is an integer,
	//if you want to see floating-point results,
	//you need to use a type coinversion to make your integers a floating point numbers.

	fmt.Printf("\n---------------------------------------------------------------------\n\n")

	var a int16 = 10
	var b int16 = 10
	var c int16 = 3
	result := a + b
	substraction := a - b
	multiplication := a * b
	division := a / c
	floatingpointdivision := float64(a) / float64(c)

	fmt.Printf(" a [%d], b[%d], c[%d] \n", a, b, c)
	fmt.Printf(" result a + b = %d\n", result)
	fmt.Printf(" substraction a - b = %d\n", substraction)
	fmt.Printf(" multiplication a * b = %d\n", multiplication)
	fmt.Printf(" division a / b = %d\n", division)
	fmt.Printf(" floating point division a / c = %f\n", floatingpointdivision)
	fmt.Printf(" floating point division ( fixed 3 decimal places) a / c = %.3f\n", floatingpointdivision)

	// +=, -=, *=, /=, %=
	// compare integers using ==, !=, >, >=, <, <=,

	fmt.Printf("\n---------------------------------------------------------------------\n\n")

	//floating point types
	//floating point literals are by default float64
	//do not use floats for exact decimal representations like money, there are other options on third party libraries for that.
	var x float32 = math.MaxFloat32
	var y float32 = -math.MaxFloat32
	var w float64 = math.MaxFloat64
	var z float64 = -math.MaxFloat64

	fmt.Printf("Minimum non-zero positive float32: %e\n", math.SmallestNonzeroFloat32)
	fmt.Printf("float32 max value: %e\n", x)
	fmt.Printf("float32 min value: %e\n", y)
	fmt.Printf("float64 max value: %e\n", w)
	fmt.Printf("float64 min value: %e\n", z)
}
