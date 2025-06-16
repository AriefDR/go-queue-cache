package main

import (
	"fmt"
)

const SIZE int = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Hash map[string]*Node

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  make(Hash),
	}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head
	return Queue{
		Head:   head,
		Tail:   tail,
		Length: 0,
	}
}

func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}

	c.Add(node)
	c.Hash[str] = node

}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Removing %s\n", n.Val)

	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.Length--
	delete(c.Hash, n.Val)
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("Add: %s\n", n.Val)

	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n
	c.Queue.Length++

	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for node != q.Tail {
		fmt.Printf("%s", node.Val)
		if node.Right != q.Tail {
			fmt.Print("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")

}

func main() {
	fmt.Println("Start Cache")

	cache := NewCache()
	for _, word := range []string{"Parrot", "Avocado", "Pineapple", "Papaya", "Peach", "Plum", "Pomegranate"} {
		cache.Check(word)
		cache.Display()
	}
}
