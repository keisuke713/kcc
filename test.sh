#!/bin/bash
assert() {
    expected="$1"
    input="$2"

    ./kcc "$input" > tmp.s
    gcc -static -g -o tmp tmp.s
    ./tmp
    actual="$?"

    if [ "$actual" = "$expected" ]; then
        echo "$input => $actual"
    else
        echo "$input => expected $expected, but got $actual"
        exit 1
    fi
}

assert 1 "1"
assert 10 "10"

// for test
// docker run --rm -v $HOME/go/src/kcc:/kcc -w /kcc compilerbook make test

// interactive
// docker run -v $HOME/go/src/kcc:/kcc -w /kcc --cap-add=SYS_PTRACE --security-opt seccomp=unconfined -it compilerbook /bin/bash