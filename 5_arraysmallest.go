package main

import "fmt"

func main() {


x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
}


var ickle int 
ickle = x[0]
fmt.Println("x 0 =  ", ickle)

for i := 0; i < len(x); i++ {
//for _, value := range x {
    if x[i] < ickle {
        ickle = x[i]
    }

}
fmt.Println("Smallest value is ", ickle)

}

