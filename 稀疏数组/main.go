package main

import (
	"fmt"
)

//稀疏数组
/*
0 0 0 0
0 1 0 0
0 0 2 0
0 0 0 0
记录上面二维数组数据，可记录非0数据的坐标（稀疏数组）
*/

type xishu struct {
	row int //行
	col int //列
	val int //值
}

func main() {

	//原始数组
	var yuanshi [11][11]int
	yuanshi[3][5] = 1
	yuanshi[6][8] = 2
	fmt.Println(yuanshi)
	//转换为稀疏数组
	//创建切片
	xs := make([]xishu, 0)
	ss := xishu{}
	for r, i := range yuanshi {
		for c, v := range i {
			if v != 0 {
				ss.row = r
				ss.col = c
				ss.val = v
				xs = append(xs, ss)
			}

		}
	}
	fmt.Println(xs)

	//稀疏数组转换为原始数组
	var ys [11][11]int
	for _, v := range xs {
		ys[v.row][v.col] = v.val
	}
	fmt.Println(ys)
}
