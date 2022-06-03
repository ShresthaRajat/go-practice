package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, To the Data Structure section")
	Numerics()
	Boolean()
}

func Numerics() {
	// There are two types of numerics:
	// int & float
	Integers()
	Float()
}

func Integers() {
	// int can be defined in following ways:
	var x int // define and assign
	x = 3
	var y int = 5 // define and assign at same
	z := 7        // use walrus operator (not reccomended)
	// (as it makes debugging harder)

	var sum = x + y + z

	fmt.Println("Sum of the numbers is: ", sum)

	// Arthmetic Operators:
	var a int = 21
	var b int = 10
	var c int

	fmt.Printf("a is %d\n", a)
	fmt.Printf("b is %d\n", b)
	c = a + b
	fmt.Printf("Line 1 a+b - Value of c is %d\n", c)

	c = a - b
	fmt.Printf("Line 2 a-b - Value of c is %d\n", c)

	c = a * b
	fmt.Printf("Line 3 a*b - Value of c is %d\n", c)

	c = a / b
	fmt.Printf("Line 4 a/b- Value of c is %d\n", c)

	c = a % b
	fmt.Printf("Line 5 a%%b - Value of c is %d\n", c)

	a++
	fmt.Printf("Line 6 a++ - Value of a is %d\n", a)

	a--
	fmt.Printf("Line 7 a-- - Value of a is %d\n", a)
}

func Float() {
	// similar to int Floats can be also defined by three different ways
	// There are two types of floats float32 and float64

	// float 32 are stored in 32 bits and stores values in single-precision floating point format
	var x float32 = 123.78
	// var y float32 = 3.4e+99 // this cannot be handled by float32

	// float 64 are stored in 64 bits and stores values in double-precision floating point format
	var z float64 = 3.4e+99
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v", z, z)
}

func Boolean() {
	// This is self explanatory
	t := true
	f := false

	fmt.Println("t != f: ", t != f)
	fmt.Println("t == f: ", t == f)
}

func String() {

}
