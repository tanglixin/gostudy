// 变量
package main

import "fmt"

func main(){
	// 声明string类型的变量
	var a string = "test"
	fmt.Println(a)

	// 声明int类型的变量
	var b,c int = 1,2
	fmt.Println(b,c)

	// 变量赋值变量，自动定义类型
	var name1 = c
	var name2 = a
	fmt.Println(name1,name2)

	//不用var声明变量， :=前面的变量要是新变量
	name3 := "lala"
	fmt.Println(name3)

	//定义各种数据类型
	var h int
	var j float64
	var i bool
	var s string
	fmt.Printf("%v %v %v %q\n",h,j,i,s)
}
