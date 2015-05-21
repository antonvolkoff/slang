package main

import (
	"fmt"
	"io"

	"github.com/k0kubun/pp"
	"github.com/peterh/liner"

	"github.com/choix/slang/s"
)

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

		output, err := s.Rep(input)
		if err != nil {
			pp.Println("error:", err)
		} else {
			pp.Println(output)
		}
	}
}
