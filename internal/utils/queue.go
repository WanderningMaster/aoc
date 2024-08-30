package utils

import "errors"

type qNode[T comparable] struct {
	next *qNode[T]
	Data T
}

type Queue[T comparable] struct {
	_map   map[T]bool
	head   *qNode[T]
	tail   *qNode[T]
	Length uint
}

func NewQueue[T comparable]() *Queue[T] {
	queue := &Queue[T]{head: nil, tail: nil, Length: 0, _map: map[T]bool{}}
	return queue
}

func (queue *Queue[T]) Enqueue(data T) {
	newNode := &qNode[T]{next: nil, Data: data}
	queue.Length += 1
	queue._map[data] = true
	if queue.head == nil {
		queue.head = newNode
		queue.tail = newNode
		return
	}
	queue.tail.next = newNode
	queue.tail = newNode
}

func (queue *Queue[T]) Dequeue() (T, error) {
	if queue.head == nil {
		var _default T
		return _default, errors.New("empty queue")
	}
	queue.Length -= 1
	data := queue.head.Data
	delete(queue._map, data)
	if queue.head == queue.tail {
		queue.head = nil
		queue.tail = nil

		return data, nil
	}
	queue.head = queue.head.next

	return data, nil
}

func (queue *Queue[T]) Peek() (T, error) {
	if queue.head == nil {
		var _default T
		return _default, errors.New("empty queue")
	}
	return queue.head.Data, nil
}

func (queue *Queue[T]) Contains(data T) bool {
	_, ok := queue._map[data]

	return ok
}
