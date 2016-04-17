 // 1
package main

import (
	"math"
	"fmt"
)

func LargestPrimeFactor1()int{
	i,s,maxPrime:=2,600851475143,0
	
	for;s!=1;{
		if(s%i==0){
			s=s/i
			maxPrime=i
		}else{
			i+=1
		}
	}
	return maxPrime
}

func LargestPrimeFactor2()int{
	n:=600851475143
	var lastFactor int
	if(n%2==0){
		n=n/2
		lastFactor=2
		for;n%2==0;{
			n=n/2
		}
	}else{
		lastFactor=1
	}
	factor:=3
	for;n>1;{
		if(n%factor==0){
			n=n/factor
			lastFactor=factor
			for;n%factor==0;{
				n=n/factor
			}
		}
		factor=factor+2
	}
	return lastFactor
	
}

func LargestPrimeFactor3()int{
	n:=600851475143
	var lastFactor int
	if(n%2==0){
		n=n/2
		lastFactor=2
		for;n%2==0;{
			n=n/2
		}
	}else{
		lastFactor=1
	}
	factor:=3
	
	maxFactor:=math.Sqrt(float64(n))
	for ;(n>1&&float64(factor)<=maxFactor);{
		if(n%factor==0){
			n=n/factor
			lastFactor=factor
			for;n%factor==0;{
				n=n/factor
			}
			maxFactor=math.Sqrt(float64(n))
		}
		factor=factor+2
	}
	if(n==1){
		return lastFactor
	}else{
		return n
	}
	
}


func main() {
	fmt.Println(LargestPrimeFactor3())
}
