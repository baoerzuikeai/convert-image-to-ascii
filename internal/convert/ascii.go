package convert

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func ConvertToASCII(input string, width uint) (string, error) {
	file, err := os.Open(input)
	if err != nil {
		return "", err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}
	height := uint(img.Bounds().Dy() * int(width) / img.Bounds().Dx() * 1 / 2)
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)
	// var grayImg image.Image
	colors := []string{
		"$", "@", "B", "%", "8", "&", "W", "M", "#", "*", "o", "a", "h", "k", "b", "d", "p", "q", "w", "m",
		"Z", "O", "0", "Q", "L", "C", "J", "U", "Y", "X", "z", "c", "v", "u", "n", "x", "r", "j", "f", "t",
		"/", "\\", "|", "(", ")", "1", "{", "}", "[", "]", "?", "-", "_", "+", "~", "<", ">", "i", "!", "l",
		"I", ";", ":", ",", "\"", "^", "`", "'", ".", " ",
	}
	var output string
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			grayColor := color.GrayModel.Convert(resizedImg.At(x, y)).(color.Gray)
			index := int(grayColor.Y) * len(colors) / 256
			output += colors[index]
		}
		output += "\n"
	}

	return output, nil
}
