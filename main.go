package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expect 2 argument")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("invalid argument")
	}

	f, _ := os.Create("tmp.s")
	fmt.Fprintf(f, ".intel_syntax noprefix\n")
	fmt.Fprintf(f, ".globl main\n")
	fmt.Fprintf(f, "main:\n")
	fmt.Fprintf(f, "  mov rax, %d\n", n)
	fmt.Fprintf(f, "  ret\n")
}
