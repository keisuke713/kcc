package main

import (
	"log"
)

type NodeKind int

const (
	NDAdd NodeKind = iota
	NDSub
	NDMul
	NDDiv
	NDNum
)

type Node struct {
	kind NodeKind
	lhs *Node
	rhs *Node
	val int
}

func newNode(kind NodeKind, lhs *Node, rhs *Node) *Node {
	return &Node{kind: kind, lhs: lhs, rhs: rhs}
}

func newNodeNum(val int) *Node {
	return &Node{kind: NDNum, val: val}
}

func expr() *Node {
	return add()
}

func add() *Node {
	node := mul()
	for {
		if consumes("+") {
			node = newNode(NDAdd, node, mul())
		} else if consumes("-") {
			node = newNode(NDSub, node, mul())
		} else {
			return node
		}
	}
}

func mul() *Node {
	node := primary()
	for {
		if consumes("*") {
			node = newNode(NDMul, node, primary())
		} else if consumes("/") {
			node = newNode(NDDiv, node, primary())
		} else {
			return node
		}
	}
}

func primary() *Node {
	if consumes("(") {
		node := expr()
		expect(")")
		return node
	}

	if token.kind == TKNum {
		return newNodeNum(expect_number())
	}

	log.Fatalf("数値でも開きカッコでもないトークンです: %s", token.str)
	return nil
}