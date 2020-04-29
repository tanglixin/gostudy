package main

import "fmt"

type Books struct{
	title string
	anchor string
	desc string
	bookid int
}

func main(){
	// 创建一个新的结构体
	fmt.Println(Books{"断舍离","李白","",1})

	// 重新赋值一个结构体
	fmt.Println(Books{title:"你当像鸟飞往你的山",anchor:"迪安",desc:"你应该像鸟一样轻,而不是一片羽毛",bookid:1})
	

}