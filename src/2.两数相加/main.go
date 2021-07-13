package main

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

func main() {
	n17 := &Node{
		Val: 9,
	}
	n16 := &Node{
		Val:  9,
		Next: n17,
	}
	n15 := &Node{
		Val:  9,
		Next: n16,
	}
	n14 := &Node{
		Val:  9,
		Next: n15,
	}
	n13 := &Node{
		Val:  9,
		Next: n14,
	}
	n12 := &Node{
		Val:  9,
		Next: n13,
	}
	n11 := &Node{
		Val:  9,
		Next: n12,
	}

	n24 := &Node{
		Val: 9,
	}
	n23 := &Node{
		Val:  9,
		Next: n24,
	}
	n22 := &Node{
		Val:  9,
		Next: n23,
	}
	n21 := &Node{
		Val:  9,
		Next: n22,
	}

	ret := combine(n11, n21)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}

func combine(n1, n2 *Node) *Node {
	ret := &Node{}
	tmp := ret
	jumpIn := 0
	var v int
	for {
		if n1 != nil && n2 != nil {
			v = n1.Val + n2.Val + jumpIn
			n1 = n1.Next
			n2 = n2.Next
		} else if n1 != nil {
			v = n1.Val + jumpIn
			n1 = n1.Next
		} else if n2 != nil {
			v = n2.Val + jumpIn
			n2 = n2.Next
		} else if jumpIn == 1 {
			v = jumpIn
		} else {
			break
		}
		if v >= 10 {
			v = v - 10
			jumpIn = 1
		} else {
			jumpIn = 0
		}
		n := &Node{}
		n.Val = v
		tmp.Next = n
		tmp = tmp.Next
	}
	return ret.Next
}
