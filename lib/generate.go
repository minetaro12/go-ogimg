package lib

import (
	"bytes"
	"strings"

	"github.com/fogleman/gg"
	"github.com/mattn/go-runewidth"
)

type SiteData struct {
	Title string
	Site  string
	Tags  []string
}

var (
	width  int = 1200
	height int = 630

	bgColor    string = "#222129"
	titleColor string = "#4169e1"
	siteColor  string = "#FFFFFF"
	tagColor   string = "#000000"
	tagBgColor string = "#FFFFFF"
)

func Generate(data SiteData) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	ctx := gg.NewContext(width, height)
	ctx.DrawRectangle(0, 0, float64(width), float64(height))
	ctx.SetHexColor(bgColor)
	ctx.Fill()

	// タイトルの書き込み
	if err := drawTitle(data.Title, ctx); err != nil {
		return nil, err
	}

	// サイト名の書き込み
	if err := drawSite(data.Site, ctx); err != nil {
		return nil, err
	}

	// タグ情報の書き込み
	if err := drawTags(data.Tags, ctx); err != nil {
		return nil, err
	}

	if err := ctx.EncodePNG(&buffer); err != nil {
		return nil, err
	}

	return &buffer, nil
}

func splitText(text string, width int) []string {
	tmp := runewidth.Wrap(text, width)
	return strings.Split(tmp, "\n")
}

func drawTitle(title string, ctx *gg.Context) error {
	if err := ctx.LoadFontFace("font.ttf", 75); err != nil {
		return err
	}

	ctx.SetHexColor(titleColor)
	tmp := splitText(title, 27)

	for k, v := range tmp {
		ctx.DrawString(v, 50, 220+float64(90*k))
	}
	return nil
}

func drawSite(site string, ctx *gg.Context) error {
	if err := ctx.LoadFontFace("font.ttf", 40); err != nil {
		return err
	}

	ctx.SetHexColor(siteColor)
	ctx.DrawString(site, 70, 570)
	return nil
}

func drawTags(tags []string, ctx *gg.Context) error {
	if err := ctx.LoadFontFace("font.ttf", 30); err != nil {
		return err
	}

	currentX := 70
	for _, v := range tags {
		ctx.SetHexColor(tagBgColor)
		// 文字列長さを取得
		width, _ := ctx.MeasureString(v)

		// タグの背景部分を描画
		ctx.DrawRectangle(float64(currentX)-10, 40, width+20, 50)
		ctx.Fill()
		ctx.SetHexColor(tagColor)
		ctx.DrawString(v, float64(currentX), 75)
		currentX = currentX + int(width) + 50
	}

	return nil
}
