package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read(input string) string {
	return input
}

func eval(ast string, env string) string {
	return ast
}

func print(exp string) string {
	return exp
}

// Read Eval Print
func rep(input string) string {
	return print(eval(read(input), ""))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("slang> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		input = strings.TrimRight(input, "\n")

		output := rep(input)
		fmt.Println(output)
	}
}
