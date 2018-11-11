package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	newHead := &Node{
		next: L.head,
		key:  key,
	}

	// Keep track of oldHead
	oldHead := L.head

	// Always insert new head
	L.head = newHead

	if oldHead == nil {
		// If oldHead was originally null, also make it the tail
		L.tail = newHead
	} else {
		// If list oldHead is not empty, add a 'prev' node entry to the oldHead
		oldHead.prev = newHead
	}

}

func (l *List) Show() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func Show(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev
	}
	fmt.Println()
}

// 1 -> 2 -> 3 -> 4 -> nil
// nil <- 1 <- 2 <- 3 <- 4
func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		curr.prev = next

		prev = curr
		curr = next

	}
	l.head = prev
	Show(l.head)
}

func main() {
	l := List{}
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)
	l.Insert(6)

	fmt.Printf("head: %v\n", l.head.key)
	fmt.Printf("tail: %v\n", l.tail.key)
	l.Show()
	l.Reverse()

	fmt.Printf("head: %v\n", l.head.key)
	fmt.Printf("tail: %v\n", l.tail.key)
}
