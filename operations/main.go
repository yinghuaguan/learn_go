package main

import (
	"fmt"
	// "math"
	"strconv"
)

// INCREMENT AND DECREMENT
// func main() {
// 	value := 10.2
// 	value++
// 	fmt.Println(value)
// 	value += 2
// 	fmt.Println(value)
// 	value -= 2
// 	fmt.Println(value)
// 	value--
// 	fmt.Println(value)
// }

// strings concatenation
// func main() {
// 	greeting := "Hello"
// 	language := "Go"
// 	combinedString := greeting + ", " + language
// 	fmt.Println(combinedString)
// }

// comparision operators
// func main() {
// 	first := 100
// 	const second = 200.00
// 	equal := first == second
// 	notEqual := first != second
// 	lessThan := first < second
// 	lessThanOrEqual := first <= second
// 	greaterThan := first > second
// 	greaterThanOrEqual := first >= second
// 	fmt.Println(equal)
// 	fmt.Println(notEqual)
// 	fmt.Println(lessThan)
// 	fmt.Println(lessThanOrEqual)
// 	fmt.Println(greaterThan)
// 	fmt.Println(greaterThanOrEqual)
// }

// comparing pointer
// func main() {
// 	first := 100
// 	second := &first
// 	third := &first
// 	alpha := 100
// 	beta := &alpha
// 	fmt.Println(second == third)
// 	fmt.Println(second == beta)
// }

// logical operators
// func main() {
// 	maxMph := 50
// 	passengerCapacity := 4
// 	airbags := true
// 	familyCar := passengerCapacity > 2 && airbags
// 	sportsCar := maxMph > 100 || passengerCapacity == 2
// 	canCategorize := !familyCar && !sportsCar
// 	fmt.Println(familyCar)
// 	fmt.Println(sportsCar)
// 	fmt.Println(canCategorize)
// }

// explicit type conversion
// func main() {
// 	kayak := 275
// 	soccerBall := 19.50
// 	total := float64(kayak) + soccerBall
// 	fmt.Println(total)
// }

// string to bool
// func main() {
// 	val1 := "true"
// 	val2 := "false"
// 	val3 := "not true"
// 	bool1, b1err := strconv.ParseBool(val1)
// 	bool2, b2err := strconv.ParseBool(val2)
// 	bool3, b3err := strconv.ParseBool(val3)
// 	fmt.Println("Bool 1", bool1, b1err)
// 	fmt.Println("Bool 2", bool2, b2err)
// 	fmt.Println("Bool 3", bool3, b3err)
// }

func main() {
	val1 := true
	val2 := false
	str1 := strconv.FormatBool(val1)
	str2 := strconv.FormatBool(val2)
	fmt.Println("Formatted value 1: " + str1)
	fmt.Println("Formatted value 2: " + str2)
}
