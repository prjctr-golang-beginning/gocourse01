package main

import "fmt"

type (
	professionName string
	operation      func(prof professionName) string
)

func (c professionName) String() string {
	return string(c)
}

func (c professionName) DisplayStatus(wp operation) {
	fmt.Println(wp(c))
}

func (c professionName) Description() string {
	switch c {
	case `nurse`:
		return "Provides care, administers medication, and assists physicians in various medical settings. There are different types of nurses, including registered nurses (RNs), nurse practitioners (NPs), and licensed practical nurses (LPNs).\n"
	case `surgeon`:
		return "Performs surgical procedures to treat or correct medical conditions.\n"
	case `dentist`:
		return "Specializes in oral health and provides services such as teeth cleaning, filling cavities, and performing dental surgeries.\n"
	case `pharmacist`:
		return "Dispenses medications, provides information on drug interactions, and advises patients on the proper use of medications.\n"
	}

	return `profession unknown`
}

const (
	office01 = iota
	office02
	office03

	ProfessionNurse   professionName = "nurse"
	ProfessionSurgeon professionName = "surgeon"
	ProfessionDentist professionName = "dentist"

	dinnerHour int = 4
	workHours  int = 9
)

func main() {
	// *** zero value
	var (
		hoursPast         int
		hasDinnerHappened bool
	)

	fmt.Printf("Zeros: %d %t\n", hoursPast, hasDinnerHappened)

	// *** small simple program
	for i := 0; i < workHours; i++ {
		if hoursPast == dinnerHour {
			hasDinnerHappened = true
			fmt.Printf("\n!!! Dinner time !!!\n\n")
		} else {
			fmt.Printf("%d work hours past\n", hoursPast)
			fmt.Printf("Has dinner happen: %t\n", hasDinnerHappened)
		}
		hoursPast++
	}

	hoursPast, hasDinnerHappened = 0, false

	// *** arguments
	printProfession(ProfessionNurse)
	printProfession(`pharmacist`)
	printProfession(`engineer`)

	fmt.Printf("%T\n", ProfessionSurgeon)
	fmt.Printf("%T\n\n", `pharmacist`)

	//// *** numbers
	arg1 := 2.2
	arg2 := -19
	fmt.Println(calc(int(arg1), arg2))

	var arg3 uint64 = 200
	var arg4 uint32 = 9
	fmt.Println(calc(int(arg3), int(arg4)))

	//// *** byte / rune
	var b1 byte = 3
	fmt.Printf("%d\n", b1)
	var b2 uint8 = 3
	fmt.Printf("%t\n", b1 == b2)

	R1 := 'Р'
	fmt.Printf("One symbol: %c, one number: %d\n", R1, R1)
	var R2 int32 = 1056
	fmt.Printf("%t\n", R1 == R2)

	smile := '😊'
	fmt.Printf("Emoji %c\n", smile)
}

func printProfession(profession professionName) {
	printOperation1 := func(prof professionName) string {
		return prof.String() +
			`: ` +
			prof.Description() +
			` works in office` + fmt.Sprintf(" %d%d%d", office03, office01, office02)
	}

	profession.DisplayStatus(printOperation1)
	profession.DisplayStatus(printOperation2)
}

func printOperation2(prof professionName) string {
	return fmt.Sprintf("%s: %s\n", prof, prof.Description())
}

func calc(a, b int) int {
	return a + b
}
