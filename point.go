//指针
package main

import "fmt"

func main(){
	var a int = 10
	var ip *int
	ip = &a
	
	fmt.Printf("a变量的地址:%x\n",&a)
	fmt.Printf("ip变量的指针:%x\n",ip)
	fmt.Printf("*ip变量的指针:%d\n",*ip)
}
