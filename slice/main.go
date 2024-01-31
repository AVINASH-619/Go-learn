package main

import (
	"fmt"
	"sort"
)

func printArrayPointer(ptr *[]int) {
	a:=int(0)
	fmt.Print(a)
	fmt.Println("Started printing")
	n := len(*(ptr))
	for i := 0; i < n; i++ {
		fmt.Println((*ptr)[i])
	}
}

func main() {
	//Using make to create a slice
	//Need to find what is the use of capacity in slice
	crr := make([]int, 2, 3) //2=>size 3=>capacity of the slice
	fmt.Printf("%T\n", crr)
	crr = append(crr, 10)
	crr = append(crr, 10) //The size of slice increase dynamically

	//[]T or []T{} or []T{values}
	brr := []int{1, 2, 4, 5, 7, 8, 9}
	//Append values to slice
	brr = append(brr, -1)
	//Len of the slice
	//fmt.Println(len(brr))
	printArrayPointer(&brr)
	//Slicing the slice
	//arr[st:end]
	sl := brr[0:4]
	fmt.Print(sl) //1 2 4 5

	//Sorting in slice

	sort.Slice(brr, func(i, j int) bool {
		return brr[i] > brr[j] //Sorts the values in desc order
	})
	for _, v := range brr {
		fmt.Println(v)
	}

	//Pointer in arr
	ptr := &brr
	fmt.Printf("%T", ptr)

}
