package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Terminal starting...")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")

		// fmt.Println(text)
		if text == "wc" {
			fmt.Println("To be implemented")
		}

	}
}
