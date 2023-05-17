package main

import (
	"fmt"
)

func main() {
	// simple array
	// declare variable
	var binatang = [4]string{"anjing", "babi", "bagong", "kambing"}
	fmt.Print(binatang)
	fmt.Println("\n")

	// akses array
	printArray(binatang)
	// rubah element pada array
	binatang[0] = "sapi"
	binatang[3] = "kebo"
	fmt.Println("array setelah dirubah \n")
	printArray(binatang)

}

// fungsi print array, disini size arraynya di set 4
func printArray(arr [4]string) {
	fmt.Println("binatang ke 1", arr[0])
	fmt.Println("binatang ke 2", arr[1])
	fmt.Println("binatang ke 3", arr[2])
	fmt.Println("binatang ke 4", arr[3], "\n")
}
