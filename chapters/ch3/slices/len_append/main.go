package main

import (
	"fmt"
	"slices"
)

// A Data structure that holds a sequence of values is a slice
// Slices can grow as needed
// the length of a slice is not part of its type itself -> this removes the limitation of the Arrays.

func main() {
	// # 1. You Do not specify the size of the slice when you declare it.
	var x = []int{10, 20, 30} //slice literal

	//you can specify indices
	var index = []int{1, 5: 10, 15: 1, 100}
	//you can simulate a multidimensional slice like a matrix
	var matrix [][]int

	// You can read and write slices using brackets syntax,
	// but you can't read or write past the end or in negative index.

	// # 2. When you do not declare a literal or values is assigned the zero value for slices -> `nil`
	// `nil` in go is a representation of lack of value for some types, nil has no type
	// a `nil` slice has nothing

	fmt.Printf("slice literal = type [%T] - Value = [%v]\n", x, x)
	fmt.Printf("index initialized = type [%T] - Value = [%v]\n", index, index)
	fmt.Printf("matrix = type [%T] - Value = [%v]\n", matrix, matrix)

	if matrix == nil {
		fmt.Println("matrix is nil")
	} else {
		fmt.Println("matrix is NOT nil")
	}

	// # 3. slices package in the standard library includes two functions to compare slices
	//3.1 `slices.Equal`, takes two slices and returns a boolean, requires to be comparable the elements of the slices
	//3.2 `slices.Func`, let you pass a function to compare, does not require the element of the slice to be comparable
	//reflect.DeepEqul is a legacy option is not safe and compare anything

	var arr1 = []int{1, 2, 3, 4, 5}
	var arr2 = []int{1, 2, 3, 4, 5}
	var arr3 = []int{1, 2, 3, 5, 7}
	fmt.Println(slices.Equal(arr1, arr2))
	fmt.Println(slices.Equal(arr1, arr3))

	// # 3 3.3 `len` -> size, passing a nil slice to len returns 0
	fmt.Printf("array size arr1 = [%d]\n", len(arr1))

	// # 3 3.3 `append` -> is used to grow slices

	var sliceEmpty []int

	sliceEmpty = append(sliceEmpty, 10)
	sliceEmpty = append(sliceEmpty, 20)
	sliceEmpty = append(sliceEmpty, 30)

	fmt.Printf("emptySlice size = [%d] - values = [%v]\n", len(sliceEmpty), sliceEmpty)

	//You can append more than one value at a time

	//Note here to reassign a slice value, has its own implications  in terms of performance so be careful
	// - Reslicing: s = s[:0] // a[low : high]
	//   - has extreme performance,  0 allocation keeps CPU warm, retain peak memory capacity in heap/stack
	//	 - has high data leak, and memory retention, keeps pointers to elements alive until overwritten
	//   - best used for High throughput recycling loops, worker pools, parsing buffers
	//  - `nil` assignment: s = nil
	//   - has Low performance, forces allocation on next append/growth
	//   - has Low memory footprint, immediately releases memory reference for GC
	//   - has no data leakage (None), because there are no references to stale data retained.
	//   - best used for short-lived task, large data pools that won't be needed for a long time
	//  - Literal Assigment: s =[]T{}
	//   A. Reslicing to zero: s = s[:0]
	// According to the specification's section on Slice expressions, a slicing operation a[low : high] creates a new slice descriptor.
	// If low is omitted, it defaults to 0. If high is 0, the resulting slice header sets Len = 0, but the underlying pointer Data remains unchanged, and Cap is preserved to the original bounds.
	//
	// B. The nil assignment: s = nil
	// Per the spec, the predeclared identifier nil represents the zero value for pointer, channel, func, interface, map, and slice types.
	// A nil slice header contains Data = 0, Len = 0, and Cap = 0
	//Under the Hood: Nil Reset (s = nil)
	//When s = nil executes, the 24-byte header is zeroed out. The reference to the backing array is dropped. If no other variables point to that backing array, the next cycle of the concurrent tri-color GC will sweep it away.
	//
	//Mechanical Sympathy Impact: When you need to read or build data again, appending to a nil slice forces runtime.growslice to hit the memory allocator (mcache/mcentral), leading to system call overheads and cache cold-starts..

	// A. Reslicing to zero: s = s[:0]
	//According to the specification's section on Slice expressions, a slicing operation a[low : high] creates a new slice descriptor.
	//If low is omitted, it defaults to 0.
	//If high is 0, the resulting slice header sets Len = 0, but the underlying pointer Data remains unchanged, and Cap is preserved to the original bounds.
	//When s = s[:0] executes, the Go compiler generates no runtime allocation calls. It simply emits machine code instructions to modify the Len field inside the stack frame's register/stack representation of the slice header.

	//Mechanical Sympathy Impact: The pointer to the memory address remains warm inside the CPU's L1/L2 data cache.
	//When you start appending new elements to s, the CPU doesn't have to wait for a hardware page fault or a runtime allocation line.
	//It writes directly to cache-resident memory.

	fmt.Printf("sliceEmpty size = [%d] - values = [%v] - capacity = [%v]\n", len(sliceEmpty), sliceEmpty, cap(sliceEmpty))
	sliceEmpty = sliceEmpty[:0]
	fmt.Printf("sliceEmpty after reslicing size = [%d] - values = [%v] - capacity = [%v]\n", len(sliceEmpty), sliceEmpty, cap(sliceEmpty))

	//The Memory Poisoning Trap (Pointer-Slice Leaks)
	/*
		type Job struct {
		    Payload [1024]byte
		}

		func ProcessBatches(jobs []*Job) {
		    // ... processing logic ...

		    // Resetting the slice for reuse
		    jobs = jobs[:0]
		}
	*/

	//If you do this, you have a massive memory leak. Even though len(jobs) is 0, cap(jobs) is still intact,
	//and the slots 0 through oldLength-1 still hold the memory addresses of your allocated Job items.
	//The Go garbage collector assumes those items are active because an reachable pointer container points to them.

	//The Architectural Fix: You must clear the elements before reslicing if they contain pointers:
	/*
		// Securely erase references to prevent GC pinned leaks
		for i := range jobs {
		    jobs[i] = nil
		}
		jobs = jobs[:0]
	*/

	//NOTE:
	//(Note: As of modern versions of Go, you can also use the standard library's clear(jobs) built-in function,
	//which zeroes out all elements of the slice while maintaining its length and capacity, allowing you to safely do clear(jobs); jobs = jobs[:0]).

}
