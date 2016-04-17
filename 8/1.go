// Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. 
package main

import (
	"strconv"	
	"bufio"
	"os"
	"fmt"
)


func test()int{
	var str string 
	var result int
	data:=make([]int,1000)
	f,err:=os.Open("data.txt")	
	defer f.Close()
	
	if err!=nil{
		panic(err)
	}
	
	scanner:=bufio.NewScanner(f)
	
	for scanner.Scan(){
		str+=scanner.Text()
	}
		
	for i:=0;i<1000;i++{	
		data[i],_=strconv.Atoi(string(str[i]))
		
	}
	
	for i:=0;i<988;i++{
		temp:=data[i]*data[i+1]*data[i+2]*data[i+3]*data[i+4]*data[i+5]*data[i+6]*data[i+7]*data[i+8]*data[i+9]*data[i+10]*data[i+11]*data[i+12]		
		if(temp>=result){
			result=temp	
		}
	}
		
	return result		
}



func main(){
	fmt.Println(test())
}
