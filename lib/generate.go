package lib

import (
	"image"
	"image/color"
	"image/draw"
	"os"

	"github.com/icza/gox/imagex/colorx"
	"golang.org/x/image/font/opentype"
)

type SiteData struct {
	Title string
	Site  string
	Tags  []string
}

const (
	width  int = 1200
	height int = 630

	bgColor    string = "#222129"
	titleColor string = "#4169e1"
	siteColor  string = "#FFFFFF"
	tagColor   string = "#000000"
	tagBgColor string = "#FFFFFF"
)

func Generate(data SiteData) (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 背景を塗る
	draw.Draw(img, img.Bounds(), &image.Uniform{hexToRGBA(bgColor)}, image.Point{}, draw.Src)

	// フォントの読み込み
	font, err := loadFont("./font.otf")
	if err != nil {
		return nil, err
	}

	drawTitle(img, data.Title, font)
	drawSite(img, data.Site, font)
	drawTags(img, data.Tags, font)

	return img, nil
}

func loadFont(path string) (*opentype.Font, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	font, err := opentype.Parse(file)
	if err != nil {
		return nil, err
	}

	return font, nil
}

func hexToRGBA(hex string) color.RGBA {
	rgba, _ := colorx.ParseHexColor(hex)
	return rgba
}
