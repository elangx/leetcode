package main

import "fmt"

func main() {
	fmt.Println(mid([]int{1, 2}, []int{3}))
}

func mid(l1, l2 []int) float64 {
	var m1, m2 int
	md := float64(len(l1)+len(l2)) / 2
	if int(md*10)%10 == 0 {
		m1 = int(md) - 1
		m2 = int(md)
	} else {
		m1 = int(md)
		m2 = m1
	}
	var i, j int
	r := 0
	for {
		t := 0
		if i < len(l1) && j < len(l2) {
			if l1[i] < l2[j] {
				t = l1[i]
				i++
			} else {
				t = l2[j]
				j++
			}
		} else if i < len(l1) {
			t = l1[i]
			i++
		} else if j < len(l2) {
			t = l2[j]
			j++
		}
		if i+j-1 == m1 {
			r += t
		}
		if i+j-1 == m2 {
			r += t
			break
		}
	}

	return float64(r) / 2
}
