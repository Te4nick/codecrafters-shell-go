package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WriteString(w *bufio.Writer, s string) {
	_, err := w.WriteString(s)
	if err != nil {
		os.Exit(1) // Exit(1) <-> IO error
	}
	err = w.Flush()
	if err != nil {
		os.Exit(1)
	}
}

func ReadString(r *bufio.Reader) string {
	s, err := r.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	running := true
	for running {
		WriteString(writer, "$ ")
		command := ReadString(reader)
		command = strings.TrimSuffix(command, "\n")
		switch command {
		case "exit":
			running = false
		default:
			WriteString(writer, fmt.Sprintf("%s: command not found\n", command))
		}
	}
}
