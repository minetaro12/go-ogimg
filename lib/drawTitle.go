package lib

import (
	"image"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func drawTitle(img *image.RGBA, title string, font_ *opentype.Font) {
	face, _ := opentype.NewFace(font_, &opentype.FaceOptions{
		Size: 75,
		DPI:  72,
	})

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(hexToRGBA(titleColor)),
		Face: face,
	}

	lineHeight := d.Face.Metrics().Height.Ceil()
	var wrappedText []string

	// 1行の長さが1100を超えたら改行する
	runes := []rune(title)
	start := 0
	for i := range runes {
		if d.MeasureString(string(runes[start:i+1])).Ceil() > 1100 {
			wrappedText = append(wrappedText, string(runes[start:i]))
			start = i
		}
	}
	// 最終行の文字列を追加
	wrappedText = append(wrappedText, string(runes[start:]))

	for lineNumber, line := range wrappedText {
		d.Dot.X = fixed.I(50)
		d.Dot.Y = fixed.I(220 + lineHeight*lineNumber)
		d.DrawString(line)
	}

}
