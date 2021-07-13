package main

import "fmt"

/**
已知自然数x、y、n， 其中x + y = n，x * 0.06四舍五入取整等于y；输入n 求x
 */

func main(){
	fmt.Println(cal(10))
}

func cal(n int)int {
	fn := float64(n)
	for i:=-0.49;i<=0.5;i+=0.01 {
		if int((fn-i)/0.0106)%100 == 0 {
			return int((fn-i)/1.06)
		}
	}
	return -1
}