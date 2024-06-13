package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type CMD func(writer *bufio.Writer, reader *bufio.Reader, args []string) error

type Handler struct {
	comMap    map[string]CMD
	reader    *bufio.Reader
	writer    *bufio.Writer
	errWriter *bufio.Writer
	path      []string
}

func NewHandler(reader *bufio.Reader, writer *bufio.Writer, errWriter *bufio.Writer, path string) *Handler {
	pathSlice := strings.Split(path, ":")
	h := &Handler{
		comMap:    make(map[string]CMD),
		reader:    reader,
		writer:    writer,
		errWriter: errWriter,
		path:      pathSlice,
	}

	h.Register("exit", h.builtinExit)
	h.Register("echo", h.builtinEcho)
	h.Register("type", h.builtinType)
	h.Register("pwd", h.builtinPwd)
	h.Register("cd", h.builtinCd)

	return h
}

func (h *Handler) Register(name string, fn CMD) {
	h.comMap[name] = fn
}

func (h *Handler) REPL() error {
	for {
		WriteString(h.writer, "$ ")
		command := ReadString(h.reader)
		command = strings.TrimSuffix(command, "\n")
		args := strings.Split(command, " ")

		cmd, ok := h.comMap[args[0]]
		if ok {
			err := cmd(h.writer, h.reader, args)
			if err != nil {
				WriteStringln(h.writer, fmt.Sprintf("%s: %s", args[0], err.Error()))
			}
			continue
		}

		cmdExt := exec.Command(args[0], args[1:]...)
		cmdExt.Stdin = os.Stdin
		cmdExt.Stdout = os.Stdout
		cmdExt.Stderr = os.Stderr

		err := cmdExt.Run()
		if err != nil {
			WriteStringln(h.writer, fmt.Sprintf("%s: command not found", args[0]))
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
	if ok {
		msg = fmt.Sprintf("%s is a shell builtin", args[1])
		WriteStringln(writer, msg)
		return nil
	}

	for _, dir := range h.path {
		fp := filepath.Join(dir, args[1])
		_, err := os.Stat(fp)
		if err == nil {
			msg = fmt.Sprintf("%s is %s", args[1], fp)
			WriteStringln(writer, msg)
			return nil
		}
	}

	msg = fmt.Sprintf("%s: not found", args[1])
	WriteStringln(writer, msg)
	return nil
}

func (h *Handler) builtinPwd(writer *bufio.Writer, _ *bufio.Reader, _ []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	WriteStringln(writer, cwd)
	return nil
}

func (h *Handler) builtinCd(_ *bufio.Writer, _ *bufio.Reader, args []string) error {
	if len(args) == 1 {
		args = append(args, os.Getenv("HOME"))
	}

	if args[1][0] == '~' {
		args[1] = strings.Replace(args[1], "~", os.Getenv("HOME"), 1)
	}

	err := os.Chdir(args[1])
	if err != nil {
		return errors.New(fmt.Sprintf("%s: No such file or directory", args[1]))
	}

	return nil
}
