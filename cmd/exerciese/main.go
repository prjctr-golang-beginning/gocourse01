package main

import "fmt"

type (
	square       float64
	printer      string
	superTrooper bool
)

func (s square) Area() float64 {
	return float64(s * s)
}

func (p printer) Println() {
	fmt.Println(p)
}

func PrintBool(val superTrooper) {
	fmt.Printf("%t\n", val)
}

func main() {
	// 1. Type, which can outputs square area
	var x square = 45.5
	fmt.Println(x.Area())

	// 2. Type, which prints itself
	var someText printer = `Hello! It's me!'`
	someText.Println()

	// 3. Func, accepted only custom type
	PrintBool(false)
	PrintBool(true)
}
