package main

import "fmt"

func main() {
	fmt.Println(cal("aba"))
}

func cal(s string) string {
	r := []rune(s)
	m := make(map[string]bool)
	for i := 0; i < len(r); i++ {
		m[fmt.Sprintf("%d-%d", i, i)] = true
	}
	c := 1
	start := 0
	for n := 2; n <= len(r); n++ {
		for i := 0; i < len(r); i++ {
			j := i + n - 1
			if j >= len(r) {
				break
			}
			if r[i] == r[j] {
				if n <= 3 {
					m[fmt.Sprintf("%d-%d", i, j)] = true
					c = n
					start = i
				} else {
					key := fmt.Sprintf("%d-%d", i, j)
					subKey := fmt.Sprintf("%d-%d", i+1, j-1)
					m[key] = m[subKey]
					if m[key] == true {
						c = n
						start = i
					}
				}
			}
		}
	}
	return string(r[start : start+c])
}
