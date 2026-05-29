package main

import (
	"fmt"
	"maps"
)

/*
Maps are a dictionary collection  where you associate a key and a value
*/
func main() {

	var nilMap map[string]int

	fmt.Printf("nilMap =[%v]\n", nilMap)
	fmt.Printf("nilMap =[%+v]\n", nilMap)

	nNilMap := map[int]int{}
	fmt.Printf("nNilMap =[%v]\n", nNilMap)
	fmt.Printf("nNilMap =[%+v]\n", nNilMap)

	nonEmptyMap := map[string][]string{
		"first":  []string{"f", "i", "r", "s", "t"},
		"second": nil,
		"third":  []string{},
	}

	fmt.Printf("nonEmptyMap =[%v]\n", nonEmptyMap)

	//you can use make to create the map

	makeMap := make(map[string]int, 10)
	fmt.Printf("makeMap =[%v]\n", makeMap)

	//reading and writing a map
	worldCupTeams := map[string]string{}
	worldCupTeams["USA"] = "NA"
	worldCupTeams["Colombia"] = "LATAM"

	fmt.Printf("worldCupTeams =[%v]\n", worldCupTeams)

	//comma, ok idiom
	//if you need to query if a key exists in a map you can use the comma ok idiom,
	//to tell if a value zero is associated with a key or is a key that is not in the map
	m := map[string]int{
		"USA":    10,
		"Canada": 5,
		"Spain":  2,
	}

	valueOk, ok := m["USA"]
	fmt.Printf("m =[%v], \"USA\" - value =[%d] - key found = [%v]\n", m, valueOk, ok)

	valueNotOk, notOk := m["Argentina"]
	fmt.Printf("m =[%v], \"Argentina\" - value =[%d] - key found = [%v]\n", m, valueNotOk, notOk)

	//Deleting from maps
	// built-in function `delete`, delete does not return a value
	delete(m, "USA")

	fmt.Printf("m =[%v]\n", m)

	//Emptying a map
	//The clear function - same as the empty ing a slice
	clear(m)
	fmt.Printf("m =[%v] - len = [%d] \n", m, len(m))

	//Comparing maps
	// Go 1.21 added a package to teh standard library called maps that contain helper functions for working with maps
	//maps.Equal and maps.EqualFunc

	letterCount := map[rune]int{'A': 10, 'B': 20, 'C': 30}
	letter2Count := map[rune]int{'A': 10, 'B': 20, 'C': 30}

	fmt.Printf("letterCount =[%c] - len = [%d]\n", letterCount, len(letterCount))
	fmt.Printf("letterCount =[%c] - len = [%d]\n", letter2Count, len(letter2Count))

	fmt.Println(maps.Equal(letterCount, letter2Count))

}
