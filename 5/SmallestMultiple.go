// SmallestMultiple
package main

import (
	"math"
	"fmt"
)

func SmallestMultiple()int{
	
	primes:=[8]int{2,3,5,7,11,13,17,19}
	//fmt.Println(primes)
	var prod int=1
	
	for i:=0;i<8;i++{
		
		var n int=2
		prime:=primes[i]
		prod=prod*prime	
		
		
		for;int(math.Pow(float64(prime),float64(n)))<21;{
			prod=prod*prime
			fmt.Println(prod)
			n+=1
		}
		
	}
	return prod
	
}

func main() {
	fmt.Println(SmallestMultiple())
}
