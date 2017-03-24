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

/*
Smallest multiple
Problem 5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
