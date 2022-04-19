package main

import "fmt"

//插入排序，将值按顺序插入相应的位置
func InsertSort(s *[6]int) {
	for i := 1; i < len(s); i++ {
		//先确定一个要插入排序的数据
		val := s[i]
		//定义插入的下标
		index := i - 1

		//确定插入位置
		//从小到大
		for index >= 0 && s[index] > val {
			//当前index的数据后移一位
			s[index+1] = s[index]
			//index下表前移一位
			index--
		}
		//插入位置为当前位置，无需再插入
		if index+1 != i {
			s[index+1] = val
		}
	}

}

func main() {
	ss := [6]int{66, 5, 77, 2, 3, 33}
	InsertSort(&ss)
	fmt.Println(ss)
}
