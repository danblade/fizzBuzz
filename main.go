package main

import "fmt"

func main() {
	for i := 0; i < 101; i++ {
		if i == 0 {
			fmt.Println(i)
		} else if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
}
