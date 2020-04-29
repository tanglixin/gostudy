package main

import "fmt"

func main(){
	var i int8 = 100
	var i1 int16 = 10000
	i2 := int32(i1) //有符号的整型
	var i3 uint = 11 //无符号的整型
	fmt.Printf("打印数据类型%T,%T\n ",i,i1) //打印类型
	fmt.Printf("二进制%b二进制%b\n",i,i1) //二进制
	fmt.Printf("十进制%d十进制%d\n",i,i1) //十进制
	fmt.Printf("八进制%o八进制%o\n",i,i1) //八进制
	fmt.Printf("十六进制%x十六进制%x\n",i,i1) //十六进制
	fmt.Printf("十六进制%x\n",i2) //十六进制
	fmt.Printf("十六进制%X\n",i3) //十六进制
}
