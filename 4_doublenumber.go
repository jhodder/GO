package main

import "fmt"

func main() {
    fmt.Print("Enter a temperature in F: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := (input - 32) * (5.0 / 9.0)
    // output := (input - 32) * 0.5555555
 
    fmt.Println(output)
}
