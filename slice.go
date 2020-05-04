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

}
