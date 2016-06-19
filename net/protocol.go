package net

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"io"
)

type request struct {
	Cmd  string
	Args [][]byte
	Conn io.ReadCloser
}

func newRequset(conn io.ReadCloser) (*request, error) {
	reader := bufio.NewReader(conn)

	// *<number of arguments>CRLF
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	var argCount int
	if line[0] == '*' {
		if _, err := fmt.Sscanf(line, "*%d\r\n", &argCount); err != nil {
			return nil, malformed("*<#Arguments>", line)
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

		return &request{
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
		return nil, malformed("$<ArgumentLenght>", line)
	}

	var argLength int
	if _, err := fmt.Sscanf(line, "$%d\r\n", &argLength); err != nil {
		return nil, malformed("$<ArgumentLength>", line)
	}

	data, err := ioutil.ReadAll(io.LimitReader(reader, int64(argLength)))
	if err != nil {
		return nil, err
	}

	if len(data) != argLength {
		return nil, malformedLength(argLength, len(data))
	}

	if b, err := reader.ReadByte(); err != nil || b != '\r' {
		return nil, malformedMissingCRLF()
	}
	if b, err := reader.ReadByte(); err != nil || b != '\n' {
		return nil, malformedMissingCRLF()
	}

	return data, nil
}

func malformed(expected string, got string) error {
	return fmt.Errorf("Mailformed request: %s does not match %s", got, expected)
}

func malformedLength(expected int, got int) error {
	return fmt.Errorf("Mailformed request: argument length %d does not match %d", got, expected)
}

func malformedMissingCRLF() error {
	return fmt.Errorf("Mailformed request: line should end with CRLF")
}

type reply io.WriterTo

type statusReply struct {
	Code string
}

func (r *statusReply) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte("+" + r.Code + "\r\n"))
	return int64(n), err
}

type errReply struct {
	Msg string
}

func (r *errReply) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte("-ERR " + r.Msg + "\r\n"))
	return int64(n), err
}

type intReply struct {
	Nos int64
}

func (r *intReply) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(":" + strconv.FormatInt(r.Nos, 10) + "\r\n"))
	return int64(n), err
}

type bulkReply struct {
	Nil  bool
	Bulk string
}

func (r *bulkReply) WriteTo(w io.Writer) (int64, error) {
	if r.Nil {
		n, err := w.Write([]byte("$-1\r\n"))
		return int64(n), err
	}
	n, err := w.Write([]byte("$" + strconv.Itoa(len(r.Bulk)) + "\r\n" + r.Bulk + "\r\n"))
	return int64(n), err
}

type listReply struct {
	Nil  bool
	List []string
}

func (r listReply) WriteTo(w io.Writer) (int64, error) {
	if r.Nil {
		n, err := w.Write([]byte("*-1\r\n"))
		return int64(n), err
	}
	if len(r.List) == 0 {
		n, err := w.Write([]byte("*0\r\n"))
		return int64(n), err
	}

	body := "*" + strconv.Itoa(len(r.List)) + "\r\n"
	for _, li := range r.List {
		if li == "" {
			body += "$-1\r\n"
			continue
		}
		body += "$" + strconv.Itoa(len(li)) + "\r\n" + li + "\r\n"
	}
	n, err := w.Write([]byte(body))
	return int64(n), err
}
