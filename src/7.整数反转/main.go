package main

import "fmt"

func main() {
	fmt.Println(rev(1230))
}
func rev(n int) int {
	output := 0
	for {
		c := n % 10
		n = n / 10
		output *= 10
		output += c
		if n == 0 {
			break
		}
	}
	return output
}
