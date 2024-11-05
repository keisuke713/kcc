package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// 文字列を数値に変換しその数字と桁数を返す
func strtoi(s string) (int, int) {
	i := 0
	j := 0
	len := len(s)
	for j < len && 48 <= s[j] && s[j] <= 57 {
		j++
	}
	n, err := strconv.Atoi(s[i:j])
	if err != nil {
		log.Fatal("invalid argument")
	}
	return n, j
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expect 2 argument")
	}

	expression := os.Args[1]
	log.Printf("%s", expression)

	f, _ := os.Create("tmp.s")
	// プロローグ
	fmt.Fprintf(f, ".intel_syntax noprefix\n")
	fmt.Fprintf(f, ".globl main\n")
	fmt.Fprintf(f, "main:\n")

	n, cnt := strtoi(expression)
	fmt.Fprintf(f, "  mov rax, %d\n", n)
	i := cnt
	len := len(expression)
	for i < len {
		switch expression[i] {
		case 43:
			i++
			n, cnt := strtoi(expression[i:])
			i += cnt
			fmt.Fprintf(f, "  add rax, %d\n", n)
		case 45:
			i++
			n, cnt := strtoi(expression[i:])
			i += cnt
			fmt.Fprintf(f, "  sub rax, %d\n", n)
		default:
			log.Fatal("unexpected character")
		}
	}
	fmt.Fprintf(f, "  ret\n")
}

// for test
// docker run --rm -v $HOME/go/src/kcc:/kcc -w /kcc compilerbook make test

// interactive
// docker run -v $HOME/go/src/kcc:/kcc -w /kcc --cap-add=SYS_PTRACE --security-opt seccomp=unconfined -it compilerbook /bin/bash
