package main

import "fmt"

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
}

func repeatLine(line string, times int) {
	for i:=0; i < times; i++ {
		fmt.Println(line)
	}
}



