package datatypes

import "fmt"

func FiddleSlice() {
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	printArray(arr)
	slc := arr[:4]
	printSlice(slc)

	slc2 := arr[4:]
	printSlice(slc2)
	//printArray(arr)
}

func printArray(a [10]int) {
	fmt.Printf("Array: %d , lenght: %d, capacity: %d\n\n", a, len(a), cap(a))
}

func printSlice(s []int) {
	fmt.Printf("Array: %d , lenght: %d, capacity: %d\n\n", s, len(s), cap(s))
}
