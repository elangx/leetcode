package main

import "fmt"

func main() {
	nums := []int64{
		2,11,15,7,
	}
	fmt.Println(getPos(nums, 9))
}
func getPos(nums []int64, target int64)(int,int) {
	mp := make(map[int64]int, len(nums))
	for i :=len(nums)-1;i>=0; i-- {
		if _, ok := mp[target-nums[i]];ok {
			return i, mp[target-nums[i]]
		}
		mp[nums[i]]=i
	}
	return -1, -1
}
