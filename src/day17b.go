package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

type nodeBuffer struct {
	nodes         []node
	nextAvailable int
}

func (buf *nodeBuffer) allocNode() *node {
	toRet := &buf.nodes[buf.nextAvailable]
	buf.nextAvailable++
	return toRet
}

func insert(cur *node, val int, steps int, nodeBuf *nodeBuffer) *node {
	for i := 0; i < steps; i++ {
		cur = cur.next
	}

	newNode := nodeBuf.allocNode()
	newNode.val = val
	newNode.next = cur.next
	cur.next = newNode

	return newNode
}

func preAllocNodes(nodeCount int) *nodeBuffer {
	var buf nodeBuffer
	buf.nodes = make([]node, nodeCount)
	return &buf
}

func main() {

	const maxVal = 50000000
	const steps = 314

	nodeBuf := preAllocNodes(maxVal + 1)

	cur := nodeBuf.allocNode()
	cur.next = cur

	for i := 1; i <= maxVal; i++ {
		cur = insert(cur, i, steps, nodeBuf)

		if i%1000000 == 0 {
			fmt.Printf("inserted val: %d million\n", i/1000000)
		}
	}

	for cur.val != 0 {
		cur = cur.next
	}

	fmt.Printf("next val: %d\n", cur.next.val)
}
