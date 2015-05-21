package main

import (
	"fmt"
	"io"

	"github.com/choix/slang/s"
	"github.com/peterh/liner"
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
			fmt.Println("error:", err)
		} else {
			fmt.Println(output)
		}
	}
}
