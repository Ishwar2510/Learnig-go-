package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	data  interface{}
	left  *TreeNode
	right *TreeNode
}
type Tree struct {
	root *TreeNode
}

func (t *Tree) Insert(data interface{}) {
	newNode := &TreeNode{data: data}
	if t.root == nil {
		t.root = newNode
		return
	}
	t.insert(data, t.root)
}

func (t *Tree) insert(data interface{}, node *TreeNode) {

	switch compare(data, node.data) {
	case 1: // we will also allow the duplicate (just for testting)
		// data > node.data  we will get 1
		if node.right == nil {
			node.right = &TreeNode{data: data}
		} else {
			t.insert(data, node.right)
		}
	case -1:
		// data < node.data
		if node.left == nil {
			node.left = &TreeNode{data: data}
		} else {
			t.insert(data, node.left)
		}
	}
}
func (t *Tree) isEmpty() error {
	if t.root == nil {
		return errors.New("Empty Tree")
	}
	return nil
}
func (t *Tree) InOrder() {
	if err := t.isEmpty(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*t)
	t.inOrder(t.root)

}
func (t *Tree) inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	t.inOrder(node.left)
	fmt.Print(node.data, " ")
	t.inOrder(node.right)
}
func (t *Tree) Preorder() {
	if err := t.isEmpty(); err != nil {
		fmt.Println(err)
		return
	}
	t.preorder(t.root)
}
func (t *Tree) preorder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.data, " ")
	t.preorder(node.left)
	t.preorder(node.right)

}
func (t *Tree) Postorder() {
	if err := t.isEmpty(); err != nil {
		fmt.Println(err)
		return
	}
	t.postorder(t.root)
}
func (t *Tree) postorder(node *TreeNode) {
	if node == nil {
		return
	}
	t.postorder(node.left)
	t.postorder(node.right)
	fmt.Print(node.data, " ")

}

func (t *Tree) Leftview() {
	myQueue := NewQueue()
	if err := t.isEmpty(); err != nil {
		fmt.Println(err)
		return
	}
	myQueue.Add(t.root)
	myQueue.PrintQueue()
	for !myQueue.IsEmpty() {
		queSize := myQueue.Size()
		for i := 1; i <= queSize; i++ {
			if i == 1 { // we need to print only the first available value for each level
				fmt.Println()
				if data, err := myQueue.Peek(); err == nil {
					// here we need to do a  type assertion as the return value is of type interface
					// and to access the field we need to do type assertion
					if res, ok := data.(*TreeNode); ok {
						fmt.Print(res.data, " ")
					}
				} else {
					fmt.Print(err)
				}
			}
			tempNode, _ := myQueue.Poll()
			if data, ok := tempNode.(*TreeNode); ok {
				if data.left != nil {
					myQueue.Add(data.left)
				}
				if data.right != nil {
					myQueue.Add(data.right)
				}
			} else {
				fmt.Print("mot ok")
			}
		}
	}
}
func compare(data1, data2 interface{}) int {
	switch data1.(type) {
	case int:
		if data1.(int) < data2.(int) {
			return -1
		}
		if data1.(int) > data2.(int) {
			return 1
		}
		return 0
	case float64:
		if data1.(float64) < data2.(float64) {
			return -1
		}
		if data1.(float64) > data2.(float64) {
			return 1
		}
		return 0
	case string:
		if data1.(string) < data2.(string) {
			return -1
		}
		if data1.(string) > data2.(string) {
			return 1
		}
		return 0

	default:
		panic("unsupported type")
	}
}

func NewTree() *Tree {
	return &Tree{}
}
