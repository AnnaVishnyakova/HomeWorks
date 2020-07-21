package main

import (
	"fmt"
	"math"
)

func main() {
	const Kurs = 65
	var Rub float64
	var Dol float64
	fmt.Println("Введите сумму в рублях:")
	fmt.Scanln(&Rub)
	Dol = math.Round(Rub / Kurs)
	fmt.Println("Сумма в долларах равна:", Dol)
}
