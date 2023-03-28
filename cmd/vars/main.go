package main

import "fmt"

type (
	color         string
	calcOperation func(a, b float64)
)

func (c color) String() string {
	return string(c)
}

const (
	ColorReset  color = "\033[0m"
	ColorBlack  color = "\u001b[30m"
	ColorRed    color = "\u001b[31m"
	ColorGreen  color = "\u001b[32m"
	ColorYellow color = "\u001b[33m"
	ColorBlue         = "\u001b[34m"

	oneWhoCannotBeNamed string = "Lord Voldemort"
)

func main() {
	// zero value
	var someZeroNumber int
	fmt.Println(someZeroNumber)

	var someFlag bool
	fmt.Println(someFlag)

	// several variables
	width, height := 100, 50
	var (
		x    = 1
		y    = 2
		form = `rectangle`
	)
	fmt.Printf("if %d * %d = %d, and this is %s\n", x, y, x*y, form)
	var pim, pom, pam string = `pim`, `pom`, `pam`
	fmt.Println(pim, pom, pam)

	// short naming
	myGender := `male`          // string
	myName, myAge := "Maks", 33 // string, int
	isCool := true              // bool
	volume := 1.1               // float

	// rune
	R := 'Р'
	fmt.Printf("One symbol: %c, one number: %d\n", R, R)

	smile := '😊'
	fmt.Printf("Emoji %c\n", smile)

	fmt.Println(`What happens with it??\n`, myGender, myName, myAge, isCool, volume)

	// functions
	textInColor(`My lvely text`, ColorBlue)
	println(`- First execution`)
	threeOperations(34, 5, Multiplication, FutureAge, Multiplication)
	println(`- Second execution`)
	threeOperations(float64(width), float64(height), Multiplication, FutureAge, Multiplication)

	// exercise
	// 1. Type, which outputs square area
	// 2. Type, which prints itself
	// 3. Func, accepted only custom type
}

func Multiplication(a, b float64) {
	fmt.Printf("%.2f\n", a*b)
}

func FutureAge(a, b float64) {
	fmt.Printf(ColorGreen.String()+"My future age is %.0f\n"+ColorReset.String(), a+b)
}

func textInColor(t string, c color) {
	fmt.Printf("%s%s\n"+ColorReset.String(), c, t)
}

func threeOperations(a, b float64, op1, op2, op3 calcOperation) {
	op1(a, b)
	op2(a, b)
	op3(a, b)
}
