package lib

import (
	"image"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func drawSite(img *image.RGBA, site string, font_ *opentype.Font) {
	face, _ := opentype.NewFace(font_, &opentype.FaceOptions{
		Size: 40,
		DPI:  72,
	})

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(hexToRGBA(siteColor)),
		Face: face,
		Dot:  fixed.Point26_6{X: fixed.I(70), Y: fixed.I(570)},
	}

	d.DrawString(site)
}
