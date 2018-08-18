package main

import (
	"fmt"
)

func main() {
	rect := Rectangle{20.5, 32.0}
	rect.Area()
	fmt.Println(rect.width)
}

type Rectangle struct {
	width  float64
	height float64
}

func (rect Rectangle) Area() {
	rect.width = 34.5
}
