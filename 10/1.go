// 1
package main

import (
	"math"
	"fmt"
)

func test()int{
	sum:=2
outer:
	for i:=3;i<2000000;i++{
		if(i%2==0){
			continue
		}else{			
			maxFactor:=int(math.Sqrt(float64(i)))
			for j:=3;j<=maxFactor;{
			if(i%j==0){
				continue outer
			}
			j+=2			
			}
			
		}
		sum+=i		
	}
	return sum
}

func main() {
	fmt.Println(test())
}
