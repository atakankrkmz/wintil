package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"

	"github.com/AllenDang/w32"
	"github.com/hnakamur/w32syscall"
)

func main() {
	err := w32syscall.EnumWindows(func(hwnd syscall.Handle, lparam uintptr) bool {
		h := w32.HWND(hwnd)
		text := w32.GetWindowText(h)
		WriteError(os.Stdout, "text => "+text+"\n")
		if strings.Contains(text, "Calculator") {
			w32.MoveWindow(h, 0, 0, 200, 600, true)
		}
		return true
	}, 0)
	if err != nil {
		WriteError(os.Stdout, err)
	}
}

func WriteError(w io.Writer, err interface{}) {
	fmt.Fprintln(w, err)
}
