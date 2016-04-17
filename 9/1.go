// 1
package main

import (
	"fmt"
)

func test()int{
	var a int
	var b int
	var c int
	for i:=1;i<1000;i++{
		for j:=1;j<=i;j++{
			a=i
			b=j
			c=1000-a-b
			if(a*a+b*b==c*c){
				return a*b*c
				break
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(test())
}
