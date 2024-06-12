package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CMD func(writer *bufio.Writer, reader *bufio.Reader, args []string) error

type Handler struct {
	comMap map[string]CMD
}

func NewHandler() *Handler {
	h := &Handler{
		comMap: make(map[string]CMD),
	}

	h.Register("exit", h.builtinExit)
	h.Register("echo", h.builtinEcho)
	h.Register("type", h.builtinType)

	return h
}

func (h *Handler) Register(name string, fn CMD) {
	h.comMap[name] = fn
}

func (h *Handler) REPL(writer *bufio.Writer, reader *bufio.Reader) error {
	for {
		WriteString(writer, "$ ")
		command := ReadString(reader)
		command = strings.TrimSuffix(command, "\n")
		args := strings.Split(command, " ")

		cmd, ok := h.comMap[args[0]]
		if !ok {
			WriteStringln(writer, fmt.Sprintf("%s: command not found", args[0]))
			continue
		}

		err := cmd(writer, reader, args)
		if err != nil {
			WriteStringln(writer, fmt.Sprintf("%s: %s", args[0], err.Error()))
		}
	}
}

func (h *Handler) builtinExit(_ *bufio.Writer, _ *bufio.Reader, args []string) error {
	if len(args) == 0 {
		return errors.New("no exit code provided")
	}

	exitCode, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("exit code must be an integer")

	}

	os.Exit(exitCode)
	return nil
}

func (h *Handler) builtinEcho(writer *bufio.Writer, _ *bufio.Reader, args []string) error {
	WriteStringln(writer, strings.Join(args[1:], " "))
	return nil
}

func (h *Handler) builtinType(writer *bufio.Writer, _ *bufio.Reader, args []string) error {
	_, ok := h.comMap[args[1]]
	var msg string
	if !ok {
		msg = fmt.Sprintf("%s: not found", args[1])
	} else {
		msg = fmt.Sprintf("%s is a shell builtin", args[1])
	}

	WriteStringln(writer, msg)
	return nil
}
