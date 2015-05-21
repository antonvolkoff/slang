package main

import (
	"fmt"
	"io"

	"github.com/k0kubun/pp"
	"github.com/peterh/liner"

	"github.com/choix/slang/s"
)

func read(input string) (*s.Node, error) {
	r := s.NewReader()
	node, err := r.Parse(input)
	// pp.Print(node)
	return node, err
}

func eval(ast *s.Node, env string) (*s.Node, error) {
	return ast, nil
}

func print(exp *s.Node) (string, error) {
	p := s.NewPrinter(exp)
	output, err := p.ToString()
	if err != nil {
		return "", err
	}
	return output, nil
}

// Read Eval Print
func rep(input string) (string, error) {
	ast, err := read(input)
	if err != nil {
		return "", err
	}

	exp, err := eval(ast, "")
	if err != nil {
		return "", err
	}

	output, err := print(exp)
	if err != nil {
		return "", err
	}

	return output, nil
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

		output, err := rep(input)
		if err != nil {
			pp.Println("error:", err)
		} else {
			pp.Println(output)
		}
	}
}
