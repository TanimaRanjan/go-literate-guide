package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	qnty := 3
	length, width := 1.2, 2.4
	len := 2
	name := "Tanima"
	fmt.Println(float64(len) * width)
	fmt.Println(qnty)
	fmt.Println(length, width)
	fmt.Println(name)

	fmt.Println("Hello, Go!!!!", "\nMy name is Tanima")
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go!!!!"))
	fmt.Println("\tCannonball!!!")
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello"))

	var originalCount int = 10
	var eatenCount int = 4

	fmt.Println("I started with", originalCount, "apples")
	fmt.Println("Some jerk ate", eatenCount, "apples")
	fmt.Println("Some jerk ate", originalCount-eatenCount, "apples left")

	price := 100
	taxRate := 0.08
	tax := float64(price) * taxRate
	fmt.Println("Tax is ", tax)
	fmt.Println(float64(price) + tax)

}
