package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// os.Argsにはコマンドライン引数が格納されている
	if len(os.Args) < 2 {
		fmt.Println("usage: ./fib [integer]")
		// 0以外を渡すとエラーを示す
		os.Exit(1)
	}
	// 文字列をint型に変換
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// 標準エラー出力
		fmt.Fprintf(os.Stderr, "argument must be integer: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d", fib(i))
}
