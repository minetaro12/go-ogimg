package lib

import (
	"bytes"
	"image"
	"image/jpeg"
	_ "image/png"
)

func Png2jpg(buf *bytes.Buffer) error {
	img, _, err := image.Decode(buf)
	if err != nil {
		return err
	}
	jpeg.Encode(buf, img, &jpeg.Options{Quality: 90})
	return nil
}
