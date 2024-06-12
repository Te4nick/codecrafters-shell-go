package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	fmt.Fprint(os.Stdout, "$ ")

	bufio.NewReader(os.Stdin).ReadString('\n')
}
