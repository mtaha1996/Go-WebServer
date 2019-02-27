package main

import "fmt"

type Node struct {
	intVal int
	strVal string
	isNone bool
}

// ---------- Stack -------------

type stack struct {
	SP      int
	Stack   []Node
	maxSize int
}

func (st *stack) push(n Node) {
	if st.SP < st.maxSize {
		st.SP++
		newStack := append(st.Stack, n)
		*st = stack{Stack: newStack}
	}
	if st.SP >= st.maxSize {
		fmt.Println("You've reached the end Stack!")
	}
}

func (st *stack) pop() Node {
	if st.SP > 0 {
		st.SP--
		return st.Stack[st.SP]
	}
	return Node{isNone: true}
}

func (st *stack) initStack(maxSize int) {
	st.SP = 0
	st.maxSize = maxSize
	st.Stack = make([]Node, maxSize)
}

func (st *stack) destroyStack() {
	st.SP = 0
	st.maxSize = 0
	st.Stack = make([]Node, 0)
}

// -------- End of Stack -------

// ---------- Queue -----------

type queue struct {
	isEmpty bool
	Queue   []Node
}

func (q *queue) enqueue(n Node) {
	newQueue := append(q.Queue, n)
	*q = queue{isEmpty: false, Queue: newQueue}
}

func (q *queue) dequeue() Node {

	// todo!
}

// -------- End of Queue -------

// -------- Linked List --------

type linkedList struct {
	head Node
	tail Node
}