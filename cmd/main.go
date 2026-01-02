package main

import (
	"flag"
	"io"
	"os"

	"github.com/baoerzuikeai/convert-image-to-ascii/internal/convert"
)

func main() {
	filepath := flag.String("p", "", "image file path")
	width := flag.Uint("w", 100, "width of output ascii art")
	mode := flag.String("m", "gray", "conversion mode:ascii (gray/clor) or hafblock")

	flag.Parse()
	switch *mode {
	case "gray", "color":
		output, err := convert.ConvertToASCII(*filepath, *width, *mode)
		if err != nil {
			// Handle error
			io.WriteString(os.Stderr, err.Error())
		}
		io.WriteString(os.Stdout, output)
	case "halfblock":
		output, err := convert.ConvertToHalfBlock(*filepath, *width, *mode)
		if err != nil {
			// Handle error
			io.WriteString(os.Stderr, err.Error())
		}
		io.WriteString(os.Stdout, output)

	default:
		io.WriteString(os.Stdout, "Unknown mode\n")
	}

	// Use the output

}
