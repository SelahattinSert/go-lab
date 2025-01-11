package main

import (
	"errors"
	"fmt"
	"math"
)

func add(x float64, y float64) float64 {
	return x + y
}

func subtract(x float64, y float64) float64 {
	return x - y
}

func multiply(x float64, y float64) float64 {
	return x * y
}

func divide(x float64, y float64) (float64, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero ")
		return 0, err
	}
	return x / y, nil
}

func factorial(x int) (int, error) {
	switch {
	case x <= 0:
		err := errors.New("x is less than zero")
		return 0, err
	case x < 50:
		i := 1
		for ; x > 0; x-- {
			i *= x
		}
		return i, nil
	}
	err := errors.New("cannot evaluate ")
	return 0, err
}

func pow(x float64, y float64) float64 {
	return math.Pow(x, y)
}

func main() {
	fmt.Println(add(23.34, 34.56))
	fmt.Println(subtract(0, 2))
	fmt.Println(multiply(9356.546456, 45654.589468))
	fmt.Println(divide(1, 0))
	fmt.Println(factorial(49))
	fmt.Println(pow(23.34, 34.56))
}
