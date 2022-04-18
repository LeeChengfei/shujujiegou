package main

import (
	"errors"
)

type NodeList struct {
	Node int
	Next *NodeList
}

//添加单链表node节点数据
func (n *NodeList) Add(newNode *NodeList) {
	//判断链表在表尾位置
	for {
		if n.Next == nil {
			break
		}
		n = n.Next
	}
	n.Next = newNode
}

//单链表有序插入
func (n *NodeList) Add2(newNode *NodeList) error {
	flag := true
	for {
		if n.Next == nil {
			break
		}
		if newNode.Node < n.Next.Node {
			break
		}
		if newNode.Node == n.Next.Node {
			flag = false
			break
		}
		n = n.Next
	}
	if !flag {
		return errors.New("有重复节点")
	} else {
		newNode.Next = n.Next
		n.Next = newNode
	}
	return nil
}

func main() {
	//定义头节点
	head := &NodeList{
		Node: 1,
	}
	n1 := &NodeList{
		Node: 2,
	}
	n2 := &NodeList{
		Node: 2,
	}
	head.Add(n1)
	head.Add(n2)

}
