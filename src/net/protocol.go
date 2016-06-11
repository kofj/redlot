package net

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"io"
)

type Request struct {
	Cmd  string
	Args [][]byte
	Conn io.ReadCloser
}

func newRequset(conn io.ReadCloser) (*Request, error) {
	reader := bufio.NewReader(conn)

	// *<number of arguments>CRLF
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	var argCount int
	if line[0] == '*' {
		if _, err := fmt.Sscanf(line, "*%d\r\n", &argCount); err != nil {
			return nil, Malformed("*<#Arguments>", line)
		}

		// $<number of bytes of argumnet 1>CRLF
		// <argument data>CRLF
		command, err := readArgumnet(reader)
		if err != nil {
			return nil, err
		}

		arguments := make([][]byte, argCount-1)
		for i := 0; i < argCount-1; i++ {
			if arguments[i], err = readArgumnet(reader); err != nil {
				return nil, err
			}
		}

		return &Request{
			Cmd:  strings.ToUpper(string(command)),
			Args: arguments,
			Conn: conn,
		}, nil
	}

	return nil, fmt.Errorf("new request error,line: %s\n", line)
}

func readArgumnet(reader *bufio.Reader) ([]byte, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, Malformed("$<ArgumentLenght>", line)
	}

	var argLength int
	if _, err := fmt.Sscanf(line, "$%d\r\n", &argLength); err != nil {
		return nil, Malformed("$<ArgumentLength>", line)
	}

	data, err := ioutil.ReadAll(io.LimitReader(reader, int64(argLength)))
	if err != nil {
		return nil, err
	}

	if len(data) != argLength {
		return nil, MalformedLength(argLength, len(data))
	}

	if b, err := reader.ReadByte(); err != nil || b != '\r' {
		return nil, MalformedMissingCRLF()
	}
	if b, err := reader.ReadByte(); err != nil || b != '\n' {
		return nil, MalformedMissingCRLF()
	}

	return data, nil
}

func Malformed(expected string, got string) error {
	return fmt.Errorf("Mailformed request: %s does not match %s", got, expected)
}

func MalformedLength(expected int, got int) error {
	return fmt.Errorf("Mailformed request: argument length %d does not match %d", got, expected)
}

func MalformedMissingCRLF() error {
	return fmt.Errorf("Mailformed request: line should end with CRLF")
}

type Reply io.WriterTo

type StatusReply struct {
	Code string
}

func (r *StatusReply) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte("+" + r.Code + "\r\n"))
	return int64(n), err
}

type ErrReply struct {
	Msg string
}

func (r *ErrReply) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte("-ERR " + r.Msg + "\r\n"))
	return int64(n), err
}

type IntReply struct {
	Nos int64
}

func (r *IntReply) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(":" + strconv.FormatInt(r.Nos, 10) + "\r\n"))
	return int64(n), err
}
