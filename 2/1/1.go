// 2
package main

import (
	"fmt"
)

func Fibonacci()int{
	a,b,fib,sum:=1,2,3,2
	for ;fib<400000000000;{
		fib=a+b
		if(fib%2==0){
			sum+=fib
		}
		a=b
		b=fib
	}
	return sum
	
}

func main() {
	fmt.Println(Fibonacci())
}
