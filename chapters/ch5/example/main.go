package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func add(a float32, b float32) (float32, error) { return a + b, nil }
func sub(a float32, b float32) (float32, error) { return a - b, nil }
func mul(a float32, b float32) (float32, error) { return a * b, nil }
func mod(a float32, b float32) (float32, error) {
	return float32(math.Mod(float64(a), float64(b))), nil
}

type mathExpression struct {
	param1    float32
	operation string
}

func main() {
	expressions := []string{
		" 2 + 3 ", " 3 - 1 ", " 10+20 ", "25+65",
	}

	//calc example using function as values
	mOperations := map[string]func(float32, float32) (float32, error){
		"+": add,
		"-": sub,
	}
	//

	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Enter your math expression")
	//
	//input, _ := reader.ReadString('\n')
	//fmt.Printf(" %s = ", input)
	for j := 0; j < len(expressions); j++ {
		tokens := strings.Split(expressions[j], " ")
		for i := 0; i < len(tokens); i++ {
			fmt.Printf("[%v]", tokens[i])
		}
		fmt.Println()
	}

	opAllowedSymbols := []byte{'+', '-', '*', '/', '%'}

	for j := 0; j < len(expressions); j++ {
		var tokenizedSt []string
		var tmp strings.Builder
		tmp.Grow(len(expressions[j]))
		currentExpression := expressions[j]
		for i := 0; i < len(currentExpression); i++ {
			if currentExpression[i] != ' ' {
				if bytes.IndexByte(opAllowedSymbols, currentExpression[i]) >= 0 {
					tokenizedSt = append(tokenizedSt, tmp.String())
					tmp.Reset()
					tmp.WriteByte(currentExpression[i])
					tokenizedSt = append(tokenizedSt, tmp.String())
					tmp.Reset()
				} else {
					tmp.WriteByte(currentExpression[i])
				}
			}
		}
		tokenizedSt = append(tokenizedSt, tmp.String())
		tmp.Reset()

		fmt.Println("TokenizedST: ", len(tokenizedSt))

		myFuncOp := mOperations[tokenizedSt[1]]
		op1, ok2 := strconv.ParseFloat(tokenizedSt[0], 32)
		op2, ok3 := strconv.ParseFloat(tokenizedSt[2], 32)

		op1f32 := float32(op1)
		op2f32 := float32(op2)
		if ok2 != nil || ok3 != nil {
			fmt.Println("there is an error parsing: ", op1, op2)
		}

		result, ok := myFuncOp(op1f32, op2f32)
		if ok == nil {
			fmt.Printf(" [ %s %v %s = %0.4f ] \n", tokenizedSt[0], tokenizedSt[1], tokenizedSt[2], result)
		}

	}
	//
	////tokenize
	//fmt.Println(" Tokenize:")
	//for j := 0; j < len(expressions); j++ {
	//	fmt.Println("expression " + expressions[j])
	//	tokenize(expressions[j], "",
	//		func(word string) bool {
	//			fmt.Printf(" -> %v \n", word)
	//			return true
	//		})
	//	fmt.Println()
	//}
}

func tokenize(s string, characters string, yield func(word string) bool) {
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '%' {
			if start < i {
				if !yield(s[start:i]) {
					return
				}
			}
			start = i + 1
		}
		if start < len(s) {
			yield(s[start:])
		}
	}
}
