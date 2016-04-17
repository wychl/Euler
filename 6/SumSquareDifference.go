// SumSquareDifference
package main

import (
	"fmt"
)

const n int=100
var m=n*(n+1)
var m2=n*(n+1)*(2*n+1)
func SumOfSquare()int{
	return m2/6
}

func SquareOfSum()int{
	
	return (m/2)*(m/2)
}

func Difference()int{
	return SquareOfSum()-SumOfSquare()
}

func main() {
	fmt.Println(Difference())
}
