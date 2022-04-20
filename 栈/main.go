package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Top int
	arr [5]int
}

func (s *Stack) Push(val int) error {
	if s.Top >= len(s.arr)-1 {
		return errors.New("栈溢出")
	}
	s.Top++
	s.arr[s.Top] = val
	return nil
}

func (s *Stack) List() error {
	if s.Top == -1 {
		return errors.New("空栈")
	}
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %d \n", i, s.arr[i])
	}
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.Top == -1 {
		return -1, errors.New("空栈")
	}
	val := s.arr[s.Top]
	s.Top--
	return val, nil
}

func main() {
	s := Stack{
		Top: -1,
	}
	for i := 0; i < 10; i++ {
		err := s.Push(i)
		if err != nil {
			break
		}
	}
	for {
		val, err := s.Pop()
		if err != nil {
			break
		}
		fmt.Println(val)
		err = s.List()
		if err != nil {
			break
		}
	}
}
