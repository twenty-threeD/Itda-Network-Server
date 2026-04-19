package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	_, _ = fmt.Fprintln(writer, "Hello World")
}
