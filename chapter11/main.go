// Reference: https://www.golang-book.com/books/intro/11

package main

import (
	"fmt"

	"example.com/chapter11/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
