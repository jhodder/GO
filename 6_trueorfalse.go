package main

import (
    "fmt"
    "bufio"
    "os"
)


func f() (int, int) {
    return 5, 6
}

func main() {

    fmt.Println("Enter an integer value: ")
    scan := bufio.NewScanner(os.Stdin)
    scan.Scan()
    fmt.Println(scan.Text())

    x, y := f()
}
