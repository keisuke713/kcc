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

# 加減算
assert 1 "1"
assert 10 "10"
assert 5 "1+2+3-1"
assert 10 "1+2+3+4"
assert 1 "2-1"
assert 1 "2+1-2"
assert 16 "10+15-9"

# 四則演算
assert 7 "1 + 2 * 3"
assert 6 "2 * 3"
assert 9 "1 + 2 * 4"
assert 4 "2 * 6 / 3"
assert 7 "2 * 3 + 1"
assert 8 "2 * ( 3 + 1 )"

echo OK
