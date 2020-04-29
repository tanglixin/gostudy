package main

import "fmt"

func main(){
	//声明数组
	var i1 [1] int

	//初始化数组
	var i2 = [5]float32{1,0.1,0.2,0.3,0.4}
	li2 := len(i2)

	//
	var i3 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	i3[4] =60.0

	//访问数组元素
	var flo50 float32 = i3[1]


	fmt.Println(i1)
	fmt.Println(i2,"长度",li2)
	fmt.Println(i3)
	fmt.Println(flo50)

	//循环数组
	for i:=0;i<len(i3);i++{
		fmt.Printf("Element[%d] = %d\n",i,i3[i])
	}
}
