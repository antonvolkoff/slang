package main

import (
	"fmt"
	"io"

	"github.com/peterh/liner"
)

func read(input string) string {
	tr := NewTokenReader(input)
	tr.Run()
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
	line := liner.NewLiner()
	defer line.Close()

	fmt.Println("Slang REPL (Ctrl-D to quit)")
	for {
		input, err := line.Prompt("slang > ")
		if err != nil {
			if err != io.EOF {
				fmt.Println("ERROR:", err)
			} else {
				fmt.Printf("\nBye Bye. See you later!\n")
			}
			return
		}

		output := rep(input)
		fmt.Println(output)
	}
}
