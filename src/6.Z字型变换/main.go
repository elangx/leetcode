package main

import "fmt"

func main() {
	fmt.Println(zSort("PAYPALISHIRING"))
}

func zSort(s string) string {
	output := ""
	r := []rune(s)
	for i := 0; i < len(r); i += 4 {
		output += string(r[i])
	}
	for i := 1; i < len(r); i += 2 {
		output += string(r[i])
	}
	for i := 2; i < len(r); i += 4 {
		output += string(r[i])
	}
	return output
}
