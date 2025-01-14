package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gibson/gophercraft/datapack"
	"github.com/gibson/gophercraft/utils/text"
)

func openFile(out string) *os.File {
	fmt.Println("Extracting to", out, "...")
	fl, _ := os.OpenFile(out, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0700)
	return fl
}

func printTimestamp(out io.Writer) {
	fmt.Fprintf(out, "// DO NOT EDIT: extracted from alpha-core database on %s\n", datapack.Timestamp())
}

type textWriter struct {
	*text.Encoder
	out io.Writer
}

func (t *textWriter) WriteComment(d string) error {
	_, err := fmt.Fprintf(t.out, "// %s\n", d)
	return err
}

func openTextWriter(out io.Writer) *textWriter {
	j := text.NewEncoder(out)
	return &textWriter{
		j, out,
	}
}
