package main

import (
	"fmt"
	"io"
)

func CountDown(out io.Writer) {
	fmt.Fprint(out, "3")
}
