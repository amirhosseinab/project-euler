package main

import "fmt"

func main() {
	const md int = 20
	fmt.Println(findSmallestEvenlyDivisible(md))
}

func findSmallestEvenlyDivisible(md int) int {
	var sd int = md
	var succeeded bool

	for {
		sd++
		succeeded = true
		for m := 1; m <= md; m++ {
			if sd%m != 0 {
				succeeded = false
				break
			}
		}
		if succeeded {
			break
		}
	}
	return sd
}
