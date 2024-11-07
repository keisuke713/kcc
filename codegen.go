package main

import (
	"io"
	"fmt"
)

func gen(w io.Writer, node *Node) {
	if node.kind == NDNum {
		fmt.Fprintf(w, "	mov rax, %d\n", node.val)
		fmt.Fprintf(w, "	push rax\n")
		return
	}

	gen(w, node.lhs)
	gen(w, node.rhs)

	fmt.Fprintf(w, "	pop rdi\n")
	fmt.Fprintf(w, "	pop rax\n")

	switch node.kind {
	case NDAdd:
		fmt.Fprintf(w, "	add rax, rdi\n")
	case NDSub:
		fmt.Fprintf(w, "	sub rax, rdi\n")
	case NDMul:
		fmt.Fprintf(w, "	imul rax, rdi\n")
	case NDDiv:
		fmt.Fprintf(w, "	cqo\n")
		fmt.Fprintf(w, "	idiv rdi\n")
	}

	fmt.Fprintf(w, "	push rax\n")
}