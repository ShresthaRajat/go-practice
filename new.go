package main

import (
    "fmt"
    )

func out(x float64, int, string) (float64, string, int){
    fmt.Println(x)
}

func main() {
    fmt.Println("Hey")
    var x int
    x = 3
    var y int = 5
    z := 7
    fmt.Println(x+y+z)
    
    if x+y+z > 10 {
    fmt.Println("big boi")
    }
    
    var a [5]int
    a[2] = 7
    fmt.Println(a)

    b := [3]int{1,2,3}
    fmt.Println(b)

//    out("Hello")


    for i := 0; i < 5; i++ {
	
	out(i)

    }

}

