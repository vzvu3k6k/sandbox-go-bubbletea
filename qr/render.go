package main

import (
	"image/color"
	"strings"

	"github.com/boombuler/barcode/qr"
)

const (
	black = "\033[40m  \033[0m"
	white = "\033[47m  \033[0m"
)

// Thanks to https://qiita.com/yoshi389111/items/fa6e592c50d2568982a4
func renderQR(data string) (string, error) {
	q, err := qr.Encode(data, qr.M, qr.Auto)
	if err != nil {
		return "", err
	}

	var out strings.Builder
	bounds := q.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if q.At(x, y) == color.Black {
				out.WriteString(black)
			} else {
				out.WriteString(white)
			}
		}
		out.WriteRune('\n')
	}

	return out.String(), nil
}
