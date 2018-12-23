package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

func insert(cur *node, val int, steps int) *node {
	for i := 0; i < steps; i++ {
		cur = cur.next
	}

	newNode := new(node)
	newNode.val = val
	newNode.next = cur.next
	cur.next = newNode

	return newNode
}

func main() {

	const maxVal = 2017
	const steps = 314

	cur := new(node)
	cur.next = cur

	for i := 1; i <= maxVal; i++ {
		cur = insert(cur, i, steps)
	}

	fmt.Printf("next val: %d\n", cur.next.val)
}
