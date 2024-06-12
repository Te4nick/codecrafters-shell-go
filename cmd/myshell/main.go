package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	fmt.Fprint(os.Stdout, "$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	switch command {
	default:
		fmt.Fprintf(os.Stdout, "%q: command not found\n", command)
	}
}
