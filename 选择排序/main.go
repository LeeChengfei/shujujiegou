package main

import "fmt"

//选择排序，从小到大排序
//先遍历找到最小的值，再进行交换
func SelectSort(s *[6]int) {
	for i := 0; i < len(s); i++ {
		//初始比较的数据和下标
		minval := s[i]
		minindex := i

		//遍历查找最小的值
		for a := i + 1; a < len(s); a++ {
			if minval > s[a] {
				minval = s[a]
				minindex = a
			}
		}

		if minindex != i {
			s[i], s[minindex] = s[minindex], s[i]
		}

	}
}

func main() {
	s := [6]int{22, 5, 1, 76, 98}
	SelectSort(&s)
	fmt.Println(s)
}
