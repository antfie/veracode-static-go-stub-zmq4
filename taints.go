package zmq4_stub

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

func stringTaint() string {
	var taint string
	fmt.Scanln(&taint)
	return taint
}

func stringArrayTaint() []string {
	taint := stringTaint()
	return strings.Split(taint, taint)
}

func byteArrayTaint() []byte {
	return []byte(stringTaint())
}

func byteArrayArrayTaint() [][]byte {
	taint := stringTaint()
	return bytes.Split([]byte(taint), []byte(taint))
}

func mapTaint() map[string]string {
	taint := stringTaint()
	var m map[string]string
	m[taint] = taint
	return m
}

func errorTaint() error {
	return errors.New(stringTaint())
}

func intTaint() int {
	i, _ := strconv.Atoi(stringTaint())
	return i
}

func int64Taint() int64 {
	return int64(intTaint())
}

func uint64Taint() uint64 {
	return uint64(intTaint())
}

func contextTaint() *Context {
	var taint string
	fmt.Scanln(&taint)

	var c = Context{
		ctx: unsafe.Pointer(&taint),
	}

	return &c
}

func socketTaint() *Socket {
	var taint string
	fmt.Scanln(&taint)

	var socket = Socket{
		soc: unsafe.Pointer(&taint),
	}

	return &socket
}
