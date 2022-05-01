package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var myArray [8]int
	fmt.Println(myArray)

	var myString [8]byte
	fmt.Println("these are strings []byte\n", myString)

	//instantiate non zero array
	zeroArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println("non zero array\n", zeroArray)

	fmt.Println(zeroArray[0])

	for i, n := range zeroArray {
		fmt.Printf("index: %d\n", i)
		fmt.Println(n)

	}
}
