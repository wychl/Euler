// LargestPalindrome
package main

import (
	"fmt"
)

func Palindrome()int{
	x:=[9]int{9,8,7,6,5,4,3,2,1}
	y:=[10]int{9,8,7,6,5,4,3,2,1,0}
	z:=[10]int{9,8,7,6,5,4,3,2,1,0}
	
	var n int
	for i:=0;i<9;i++{
		for j:=0;i<10;j++{
			for k:=0;k<10;k++{
				n=100001*x[i]+10010*y[j]+1100*z[k]
				for l:=999;l>=100;l--{
					if(n%l==0){
						n2:=n/l
						if(100<=n2&&n2<=999){
							return l*n2
						}else{
							continue
						}
						
					}
				}
			}
		}
	}
	return 0
	
}

func main() {
	fmt.Println(Palindrome())
}
