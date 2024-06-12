package main

import (
	"bufio"
	"os"
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

	h := NewHandler()

	err := h.REPL(writer, reader)
	if err != nil {
		panic(err)
	}
}
