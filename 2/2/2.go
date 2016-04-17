// 2
package main

import (
	"fmt"
)

func Fibonacci()int{
	x,y,sum:=1,1,0
	for ;sum<40000000000;{
		sum+=x+y
		x,y=x+2*y,2*x+3*y
	}
	return sum
}

func main() {
	fmt.Println(Fibonacci())
}
