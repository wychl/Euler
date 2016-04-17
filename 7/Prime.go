// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13
// we can see that the 6th prime is 13.
//What is the 10 001st prime number?
package main

import (
	//"math"
	"fmt"
)

func Prime()int{
	var count int=1
	var i int=3
	
outer:
	for ;;i++{		
		if(i%2==0){
			continue
		}else{						
			for j:=2;j<i;j++{				
				if(i%j==0){					
					continue outer
				}			
			}
			count+=1			
		}
											
		if(count==10001){
			return i
		}	
			
		}
}

func main() {
	fmt.Println(Prime())
}
