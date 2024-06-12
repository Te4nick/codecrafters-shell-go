package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func WriteStringln(w *bufio.Writer, s string) {
	WriteString(w, s+"\r\n")
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

	for {
		WriteString(writer, "$ ")
		command := ReadString(reader)
		command = strings.TrimSuffix(command, "\n")
		args := strings.Split(command, " ")
		switch args[0] {
		case "exit":
			exitCode, err := strconv.Atoi(args[1])
			if err != nil {
				WriteStringln(writer, "error: argument must be an integer")
				break
			}
			os.Exit(exitCode)
		default:
			WriteStringln(writer, fmt.Sprintf("%s: command not found", args[0]))
		}
	}
}
