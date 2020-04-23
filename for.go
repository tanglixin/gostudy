//冒泡排序
package main

import "fmt"

func main(){
	var list = []int{ 2,1,4,7,6}
	fmt.Println(list)
	BubbleSort(list)
	BubbleZSort(list)
}

//生序
func BubbleSort(list []int){
	len := len(list)
	for i := 0; i < len; i++{
		for j := i+1; j < len; j++{
			if list[i] > list[j] {
				list[i],list[j] = list[j],list[i]
			}
		}
	}

	fmt.Println(list)
}

//降序
func BubbleZSort(list []int){
	len := len(list)
	for i := 0; i < len; i++{
		for j := i+1; j < len; j++{
			if list[i] < list[j] {
				list[i],list[j] = list[j],list[i]
			}
		}
	}

	fmt.Println(list)
}
