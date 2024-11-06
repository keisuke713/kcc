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
	log.Printf("%s", expression)

	token = tokenize(expression)

	// アセンブリ前半部分を出力
	f, _ := os.Create("tmp.s")
	// プロローグ
	fmt.Fprintf(f, ".intel_syntax noprefix\n")
	fmt.Fprintf(f, ".globl main\n")
	fmt.Fprintf(f, "main:\n")

	// 式の最初は数でなければならないので、それをチェックして
	// 最初のmov命令を出力
	fmt.Fprintf(f, "	mov rax, %d\n", expect_number())

  	// `+ <数>`あるいは`- <数>`というトークンの並びを消費しつつ
  	// アセンブリを出力
	for !at_eof() {
		if consumes("+") {
			fmt.Fprintf(f, "	add rax, %d\n", expect_number())
			continue
		}

		expect("-")
		fmt.Fprintf(f, "	sub rax, %d\n", expect_number())
	}
	fmt.Fprintf(f, "  ret\n")
}

// for test
// docker run --rm -v $HOME/go/src/kcc:/kcc -w /kcc compilerbook make test

// interactive
// docker run -v $HOME/go/src/kcc:/kcc -w /kcc --cap-add=SYS_PTRACE --security-opt seccomp=unconfined -it compilerbook /bin/bash
