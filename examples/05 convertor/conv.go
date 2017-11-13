package main

import "fmt"

const (
	miToKm  = 1.60934
	pToKg   = 0.453592
	divider = "+-----------------+"
)

func main() {
	fmt.Println("Choose an option")
	fmt.Println(`1. Miles to Km
2.farenheit to celsius
3.Pound to kg
		`)

	var (
		option int
		number float64
	)
	fmt.Scanln(&option)
	fmt.Println("Enter the value")
	fmt.Scanln(&number)

	switch option {
	case 1:

		fmt.Printf("|Miles: %10.2f |\n", number)
		fmt.Println(divider)
		fmt.Printf("|Kilometers: %10.2f |\n", number*miToKm)

	case 2:
		fmt.Printf("|Farenheit: %10.2f |\n", number)
		fmt.Println(divider)
		fmt.Printf("|Celsius: %10.2f |\n", (number-32)*5/9)

	case 3:
		fmt.Printf("|Pound: %10.2f |\n", number)
		fmt.Println(divider)
		fmt.Printf("|Kilograms: %10.2f |\n", number*pToKg)

	}

}
