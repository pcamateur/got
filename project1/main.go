package main

import "fmt"

func main() {
	var x, y int

	for x = 0; x < 10; x++ {
		for y = 1; y < x; y++ {
			if x%y == 2 {
				fmt.Println(x, y)
			}
		}
	}

}
