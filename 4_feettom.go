package main

import "fmt"

func main() {

    const scalefactor float64  = 0.3048

    fmt.Print("Enter length in Feet: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := (input * scalefactor)
    // output := (input - 32) * 0.5555555
 
    fmt.Println(output)
}
