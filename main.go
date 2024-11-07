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
	node := expr()

	// アセンブリ前半部分を出力
	f, _ := os.Create("tmp.s")
	// プロローグ
	fmt.Fprintf(f, ".intel_syntax noprefix\n")
	fmt.Fprintf(f, ".globl main\n")
	fmt.Fprintf(f, "main:\n")

	gen(f, node)
	// 最終的な戻り値をraxにセット
	fmt.Fprintf(f, "	pop rax\n")
	fmt.Fprintf(f, "	ret\n")
}

// for test
// docker run --rm -v $HOME/go/src/kcc:/kcc -w /kcc compilerbook make test

// interactive
// docker run -v $HOME/go/src/kcc:/kcc -w /kcc --cap-add=SYS_PTRACE --security-opt seccomp=unconfined -it compilerbook /bin/bash
