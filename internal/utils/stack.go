package utils

import "errors"

type sNode struct {
	next *sNode
	Data interface{}
}

type Stack struct {
	head   *sNode
	Length uint
}

func NewStack() *Stack {
	stack := &Stack{head: nil, Length: 0}
	return stack
}

func (stack *Stack) Push(data interface{}) {
	newNode := &sNode{next: nil, Data: data}
	stack.Length += 1
	if stack.head == nil {
		stack.head = newNode
		return
	}
	head := stack.head
	stack.head = newNode
	newNode.next = head
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.head == nil {
		return nil, errors.New("empty stack")
	}
	stack.Length -= 1
	data := stack.head.Data
	stack.head = stack.head.next

	return data, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.head == nil {
		return nil, errors.New("empty stack")
	}
	return stack.head.Data, nil
}
