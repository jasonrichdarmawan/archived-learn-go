// Variables

package main

import (
	"fmt"
	math "main/math"
)

func main() {
	var x string = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	var z string
	z = "first"
	fmt.Println(z)
	z = "second"
	fmt.Println(z)

	var a string
	a = "first "
	fmt.Println(a)
	a = a + "second"
	fmt.Println(a)

	var b string = "hello"
	var c string = "world"
	fmt.Println(b == c) // equality

	d := "Hello, World" // shorter statement to create a new variable
	fmt.Println(d)

	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

	fmt.Println(math.FahrenheitToCelsius(20)) // Write a program that converts from Fahrenheit into Celsius

	fmt.Println(math.FeetToMeters(2)) // Write another program that converts from feet into meters (1 ft = 0.3048 m)
}
