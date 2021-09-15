package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string // vazia por padrão -> %q irá mostrar com ""
	var e *int

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}
