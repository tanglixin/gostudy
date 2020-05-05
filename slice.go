//切片
package main

import "fmt"

func main(){

	//定义一个存放int类型的切片，没有限制长度，底层是一个数组
	var i [] int
	//定义一个string类型的切片
	var s [] string

	i = []int{1,2,3}
	s = []string{"北京","河北","天津"}

	fmt.Println(i,s)

	//len获取长度、cap可以测量切片最大能达到长度
	fmt.Printf("len(i)=%d,cap(i)=%d \n",len(i),cap(i))
	fmt.Printf("len(s)=%d,cap(s)=%d \n",len(s),cap(s))


	//由数组得到切片,基于一个数组切割，左包含右不包含
	a1 := [...]int{1,3,5,7,9,11,13}
	a3 := a1[0:4]
	a4 := a1[:4]
	a5 := a1[:]
	a6 := a1[1:5]
	fmt.Println(a1,a3,a4,a5)

	//切片的容量是指底层数组的容量
	fmt.Println(len(a3),cap(a3)) //4,7
	fmt.Println(len(a6),cap(a6)) //4,6


	//切片再切片,指向底层的数组
	a6[1] = 1000
	a7 := a6[0:]
	fmt.Println(a7,len(a7),cap(a7))


	var numbers []int
	//追加空切片
	numbers = append(numbers,0) //[0]
	//向切片添加一个元素
	numbers = append(numbers, 1)//[0 1]
	//同时添加多个元素
	numbers = append(numbers, 2,3,4)//[0 1 2 3 4]
	fmt.Println(numbers)

	//创建切片 numbers1 是之前切片的两倍容量
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	//拷贝 numbers 的内容到 numbers1
	copy(numbers1,numbers)
	fmt.Println(numbers1)

}
