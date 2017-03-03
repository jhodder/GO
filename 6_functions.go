package main

import "fmt"

func main() {
    slice1 := []int{1,2,3}
    slice2 := make([]int, 2)
    slice3 := make([]int, 3, 9)
    x := [6]string{"a","b","c","d","e","f"}

    fmt.Println("Before")
    fmt.Println(slice1, slice2)

    fmt.Println("After")
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
    fmt.Println("this")
    fmt.Println(len(slice3))
    fmt.Println("Array at 2:5 is ", x[2:5])

}
