package lib

import (
	"image"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func drawSite(img *image.RGBA, site string, font_ *truetype.Font) {
	face := truetype.NewFace(font_, &truetype.Options{
		Size: 40,
	})

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(hexToRGBA(siteColor)),
		Face: face,
		Dot:  fixed.Point26_6{X: fixed.I(70), Y: fixed.I(570)},
	}

	d.DrawString(site)
}
