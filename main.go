package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expect 2 argument")
	}

	expression := os.Args[1]

	token = tokenize(expression)
	nodes := program()

	// アセンブリ前半部分を出力
	f, _ := os.Create("tmp.s")
	fmt.Fprintf(f, ".intel_syntax noprefix\n")
	fmt.Fprintf(f, ".globl main\n")
	fmt.Fprintf(f, "main:\n")
	// プロローグ
	fmt.Fprintf(f, "	push rbp\n")
	fmt.Fprintf(f, "	mov rbp, rsp\n")
	// 変数 `a` ~ `z` 分のスタック領域を確保
	fmt.Fprintf(f, "	sub rsp, 208\n")

	for _, n := range nodes {
		if n == nil {
			break
		}
		gen(f, n)
		// 戻り値をraxにセット
		fmt.Fprintf(f, "	pop rax\n")
	}
	// エピローグ
	fmt.Fprintf(f, "	mov rsp, rbp\n")
	fmt.Fprintf(f, "	pop rbp\n")
	fmt.Fprintf(f, "	ret\n")
}

// for test
// docker run --rm -v $HOME/go/src/kcc:/kcc -w /kcc compilerbook make test

// interactive
// docker run -v $HOME/go/src/kcc:/kcc -w /kcc --cap-add=SYS_PTRACE --security-opt seccomp=unconfined -it compilerbook /bin/bash
