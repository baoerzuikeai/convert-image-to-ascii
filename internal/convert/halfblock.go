package convert

import (
	"image"
	"os"
	"strconv"

	"github.com/nfnt/resize"
)

func ConvertToHalfBlock(input string, width uint, mode string) (string, error) {
	file, err := os.Open(input)
	if err != nil {
		return "", err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}
	height := uint(img.Bounds().Dy() * int(width) / img.Bounds().Dx())
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	var output string
	for y := 0; y < int(height); y += 2 {
		for x := 0; x < int(width); x++ {
			r1, g1, b1, _ := resizedImg.At(x, y).RGBA()
			strr1 := strconv.Itoa(int(r1 >> 8))
			strg1 := strconv.Itoa(int(g1 >> 8))
			strb1 := strconv.Itoa(int(b1 >> 8))
			var strr2, strg2, strb2 string
			if y+1 < int(height) {
				r2, g2, b2, _ := resizedImg.At(x, y+1).RGBA()
				strr2 = strconv.Itoa(int(r2 >> 8))
				strg2 = strconv.Itoa(int(g2 >> 8))
				strb2 = strconv.Itoa(int(b2 >> 8))
			} else {
				strr2 = "0"
				strg2 = "0"
				strb2 = "0"
			}

			output += "\x1b[;2;" + strr1 + ";" + strg1 + ";" + strb1 + "m" + "\x1b[48;2;" + strr2 + ";" + strg2 + ";" + strb2 + "m" + "▀" + "\x1b[0m"

		}
		output += "\n"
	}

	return output, nil
}
