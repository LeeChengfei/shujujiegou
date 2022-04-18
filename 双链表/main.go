package main

import (
	"errors"
	"fmt"
)

type NodeList struct {
	Node int
	Last *NodeList
	Next *NodeList
}

//链表尾部添加节点
func (n *NodeList) Add(newNode *NodeList) {
	for {
		if n.Next == nil {
			break
		}
		n = n.Next
	}
	newNode.Last = n
	n.Next = newNode
}

//删除指定节点
func (n *NodeList) Remove(id int) error {
	flag := false
	//找到要删除的节点
	for {
		if n.Next == nil {
			break
		}
		if n.Next.Node == id {
			flag = true
			break
		}
		n = n.Next
	}
	if flag {
		n.Next = n.Next.Next
		//预防要删除的节点是最后的一个节点
		//因为最后一个节点的Next是空
		if n.Next != nil {
			n.Next.Last = n
		}
	} else {
		return errors.New("没有找到要删除的节点")
	}
	return nil
}

func main() {
	head := &NodeList{
		Node: 1,
	}
	n1 := &NodeList{
		Node: 2,
	}
	n2 := &NodeList{
		Node: 3,
	}
	head.Add(n1)
	head.Add(n2)
	err := head.Remove(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(head.Next.Node)
}
