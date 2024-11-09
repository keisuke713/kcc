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
	NDAssign
	NDLVar
)

type Node struct {
	kind NodeKind
	lhs *Node
	rhs *Node
	val int
	offset int
}

func newNode(kind NodeKind, lhs *Node, rhs *Node) *Node {
	return &Node{kind: kind, lhs: lhs, rhs: rhs}
}

func newNodeNum(val int) *Node {
	return &Node{kind: NDNum, val: val}
}

// program = stmt*
func program() []*Node {
	nodes := make([]*Node, 0, 10)
	for token.kind != TKEOF {
		nodes = append(nodes, stmt())
	}
	return nodes
}

// stmt = expr ";"
func stmt() *Node {
	node := expr()
	expect(";")
	return node
}

// expr = assign
func expr() *Node {
	return assign()
}

// assign = add ("=" add)?
func assign() *Node {
	node := add()
	if consumes("=") {
		node = newNode(NDAssign, node, assign())
	}
	return node
}

// add = mul ("+" mul | "-" mul)*
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

// mul = primary ("*" primary | "/" primary)*
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

// primary = num | "(" expr ")" | ident
func primary() *Node {
	if token.kind == TKIdent {
		node := newNode(NDLVar, nil, nil)
		node.offset = int(token.str[0] - 96) * 8
		token = token.next
		return node
	}

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