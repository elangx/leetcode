package main

import "fmt"

func main() {
	fmt.Println(longNoneRepeat(""))
}

func longNoneRepeat(s string) int {
	var c int
	var ret int
	m := make(map[rune]bool)
	for _, v := range s {
		if _, ok := m[v]; ok {
			m = make(map[rune]bool)
			if ret < c {
				ret = c
			}
			c = 1
		} else {
			c++
		}
		m[v] = true
	}
	return ret
}
