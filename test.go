package main

import (
	"fmt"
	"reflect"
)

func test() {
	var intSlice = make([]int, 10)

	fmt.Printf("intSlice: \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(reflect.ValueOf(intSlice).Kind())

	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 2

	fmt.Println("Slice A:", a)
	fmt.Printf("Length is %d capacity is %d\n", len(a), cap(a))

	a = append(a, 30, 18, 99, 55)
	fmt.Println("Slice A after append:", a)
	fmt.Printf("Length is %d capacity is %d", len(a), cap(a))
}
