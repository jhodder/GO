package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
	var fizz bool = false
	var buzz bool = false
	if i % 3 == 0 {
	  fizz = true
	}
	if i % 5 == 0 {
	  buzz = true
	}
	fmt.Println(i, fizz, buzz)
	
    }
}
