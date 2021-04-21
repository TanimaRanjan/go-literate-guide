package main

import (
	"fmt"
	"math"
)

var metersPerLiter float64

func paintMe(width float64, height float64) float64 {
	area := width * height
	return area / metersPerLiter
}

func main() {
	metersPerLiter = 10.0
	repeatLine("hello", 3)
	fmt.Printf("%.2f \n", paintMe(4.2, 3.0))
	fmt.Printf("%.2f \n", paintMe(5.2, 3.0))
	fmt.Printf("%.2f \n", paintMe(5.2, 3.3))
	cans, remainder := floatParts(1.26)
	fmt.Println(cans, remainder)
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f liters needed \n", amount)
	}
}

func repeatLine(line string, times int) {
	for i:=0; i < times; i++ {
		fmt.Println(line)
	}
}

func floatParts(number float64) (integerPart int, fractionPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 { 
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil

}

