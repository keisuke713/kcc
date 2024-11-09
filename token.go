package main

import (
	"log"
	"strconv"
)

// token type
type TokenKind int

const (
    TKReserved TokenKind = iota
    TKNum
	TKIdent
    TKEOF
)

type Token struct {
	kind TokenKind
	next *Token // 次の入力のトークン
	val int // kindがTKNumの場合数値が入る
	str string // トークン文字列
}

// 現在着目しているトークン
var token *Token

// 次のトークンが期待している記号のときはトークンを一つ読み進めて
// 真を返す。それ以外の場合は偽を返す。
func consumes(op string) bool {
	if (token.kind != TKReserved) || (len(op) != len(token.str)) || (op != token.str) {
		return false
	}
	token = token.next
	return true
}

// 次のトークンが期待している記号のときはトークンを一つ読み進める
// それ以外の場合はエラーを報告する
func expect(op string) {
	if (token.kind != TKReserved) || (len(op) != len(token.str)) || (op != token.str) {
		log.Fatalf("'%s'ではありません", op)
	}
	token = token.next
}

// 次のトークンが数値の場合、トークンを一つ読み進めてその数値を返す
// それ以外の場合はエラーを報告する
func expect_number() int {
	if token.kind != TKNum {
		log.Fatal("数ではありません")
	}
	val := token.val
	token = token.next
	return val
}

func at_eof() bool {
	return token.kind == TKEOF
}

// 新しいトークンを作成してcurに繋げる
func newToken(kind TokenKind, cur *Token, str string) *Token {
	tok := &Token{kind: kind, str: str}
	cur.next = tok
	return tok
}

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

// 入力文字列strをトークナイズしてそれを返す
func tokenize(str string) *Token {
	p := 0
	head := &Token{}
	cur := head

	for p < len(str) {
		if str[p] == ' ' {
			p++
			continue
		}

		if (str[p] == '+') || (str[p] == '-' || (str[p] == '*') || (str[p] == '/') || (str[p] == '(') || (str[p] == ')') || (str[p] == '=') || (str[p] == ';')) {
			cur = newToken(TKReserved, cur, str[p:p+1])
			p++
			continue
		}

		if ('0' <= str[p]) && (str[p] <= '9') {
			cur = newToken(TKNum, cur, "")
			// cur.val = int(str[p] - '0')
			// p++
			n, cnt := strtoi(str[p:])
			cur.val = n
			p += cnt
			continue
		}

		if 'a' <= str[p] && str[p] <= 'z' {
			cur = newToken(TKIdent, cur, str[p:p+1])
			p++
			continue
		}

		log.Fatal("トークナイズできません")
	}

	newToken(TKEOF, cur, "")
	return head.next
}
