package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "hiiii")
	io.WriteString(os.Stdout, "hiiii")

}
