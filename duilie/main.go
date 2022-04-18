package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Size  int    //队列大小
	Array [4]int //数组实现队列操作
	Front int    //队首
	Rear  int    //队尾
}

//队列从队尾添加数据
func (q *Queue) Add(val int) error {
	if q.Rear == q.Size-1 {
		return errors.New("队列已满")
	}
	q.Rear++ //队尾后移一位
	q.Array[q.Rear] = val
	return nil
}

//队列从队首获取数据
func (q *Queue) Get() (int, error) {
	//判断队列是否到队尾
	if q.Front == q.Rear {
		return 0, errors.New("队列已空")
	}
	//初始化队列时定义队首为-1，这里需要先自加再取值
	q.Front++
	val := q.Array[q.Front]
	return val, nil
}
func main() {
	//初始化一个队列
	q := &Queue{
		Size:  4,
		Front: -1,
		Rear:  -1,
	}

	//向队列添加数据
	for i := 0; i < 5; i++ {
		err := q.Add(i)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	//读取队列数据
	for {
		val, err := q.Get()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(val)
	}
}
