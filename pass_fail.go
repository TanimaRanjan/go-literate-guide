package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
//	file, _ := os.Stat("hello.go")
//	fmt.Println(file.Size())

	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, _ := strconv.ParseFloat(input, 64)

	var status string
	if err == nil {
		if grade == 100 {
			status = "Perfect!"
		} else if grade >= 60 {
			status = "You pass."
		} else {
			status ="You fail!"
		}
	}
	fmt.Println("A grade of", grade, "is", status)
}
