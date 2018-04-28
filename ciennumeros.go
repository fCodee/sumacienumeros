package main

import "fmt"

func main() {
	var temp int
	for i := 1; i <= 100; i++ {
		temp += i
		fmt.Println(i)
	}

	fmt.Println("la suma de los numeros es ", temp)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}
