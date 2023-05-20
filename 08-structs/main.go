package main

import (
	"fmt"
)

type Rectangle struct {
	Width  float64
	Height float64
}


func main() {
	rect := Rectangle{Width: 5, Height: 10}
	area := rect.Area()
	fmt.Println("Area:", area)
}
// Method defined on the Rectangle struct
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}