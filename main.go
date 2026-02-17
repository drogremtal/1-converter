package main

import "fmt"

func main() {
	const USDTOEURO float64 = 0.88
	const USDTORUBL float64 = 72.77

	///кол-во евро
	var number float64 = 45.66

	fmt.Println("Конвертация USD:", number)
	fmt.Println(number * (1 / USDTOEURO) * USDTORUBL)

}
