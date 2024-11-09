package main

import (
	"io"
	"fmt"
	"log"
)

func genLeftVal(w io.Writer, node *Node) {
	if node.kind != NDLVar {
		log.Fatalf("代入の左辺値が変数ではありません")
	}
	fmt.Fprintf(w, "	mov rax, rbp\n")
	fmt.Fprintf(w, "	sub rax, %d\n", node.offset)
	fmt.Fprintf(w, "	push rax\n")
}

func gen(w io.Writer, node *Node) {
	if node == nil {
		return
	}

	switch node.kind {
	case NDNum:
		fmt.Fprintf(w, "	mov rax, %d\n", node.val)
		fmt.Fprintf(w, "	push rax\n")
		return
	case NDAssign:
		genLeftVal(w, node.lhs)
		gen(w, node.rhs)

		fmt.Fprintf(w, "	pop rdi\n")
		fmt.Fprintf(w, "	pop rax\n")
		fmt.Fprintf(w, "	mov [rax], rdi\n")
		fmt.Fprintf(w, "	push rdi\n")
		return
	case NDLVar:
		genLeftVal(w, node)

		fmt.Fprintf(w, "	pop rax\n")
		fmt.Fprintf(w, "	mov rax, [rax]\n")
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