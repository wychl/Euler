// SumDivisibleBy
package main

import (
	"fmt"
)
const target int=99999
func SumDivisibleBy(n int)int{
	p:=target/n
	return n*(p*(p+1))/2
}


func main() {
	fmt.Println(SumDivisibleBy(3)+SumDivisibleBy(5)-SumDivisibleBy(15))
}
