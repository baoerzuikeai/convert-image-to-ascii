package main

import (
	"io"
	"os"

	"github.com/baoerzuikeai/convert-image-to-ascii/internal/convert"
)

func main() {
	output, err := convert.ConvertToASCII("./test.jpg", 150)
	if err != nil {
		// Handle error
	}
	// Use the output
	io.WriteString(os.Stdout, output)
}
