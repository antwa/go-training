package main

import "fmt"

func main() {
	// memanggil fungsi
	printMsg()
	// pemanggilan fungsi yang memiliki return value dan memasukanya ke variable
	var myMessage string = getMessage()
	fmt.Println(myMessage)

	// memangil fungsi yang memiliki return value dan memiliki parameter dan memasukanya ke variable
	var panjang float32 = 50
	var lebar float32 = 20
	var luasPersegi float32 = hitungLuasPersegi(panjang, lebar)
	fmt.Println("Luas Persegi dari dengan variable:")
	fmt.Println("Panjang :", panjang)
	fmt.Println("Lebar :", lebar)
	fmt.Println("Luas Persegi :", luasPersegi)

}

// fungsi biasa
func printMsg() {
	fmt.Println("call from printMsg")
}

// fungsi keluaran
func getMessage() string {
	return "ini hasil keluaran dari fungsi"
}

// fungsi keluaran menggunakan argument/parameter
func hitungLuasPersegi(panjang float32, lebar float32) float32 {
	// panjang tipe datanya float32
	// lebar tipe datanya float32
	var luas float32 = 0
	luas = panjang * lebar
	return luas
}
