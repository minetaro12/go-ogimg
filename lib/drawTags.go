package lib

import (
	"image"
	"image/draw"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func drawTags(img *image.RGBA, tags []string, font_ *truetype.Font) {
	face := truetype.NewFace(font_, &truetype.Options{
		Size: 30,
	})

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(hexToRGBA(tagColor)),
		Face: face,
	}

	currentX := 70
	for _, tag := range tags {
		// タグの長さを取得
		width := d.MeasureString(tag).Ceil()

		// タグの背景を描画
		draw.Draw(img, image.Rect(currentX-10, 40, currentX+width+10, 90), &image.Uniform{hexToRGBA(tagBgColor)}, image.Point{}, draw.Src)

		// タグを描画
		d.Dot.X = fixed.I(currentX)
		d.Dot.Y = fixed.I(75)
		d.DrawString(tag)

		currentX = currentX + width + 50
	}
}
