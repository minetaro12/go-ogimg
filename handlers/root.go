package handlers

import (
	"bytes"
	"fmt"
	"go-ogimg/lib"
	"image/jpeg"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Root(c *fiber.Ctx) error {
	// クエリ文字列のタグ部分のパース
	tags := []string{}
	if c.Query("tags") != "" {
		tags = strings.Split(c.Query("tags"), ",")
	}

	data := lib.SiteData{
		Title: c.Query("title"),
		Site:  c.Query("site"),
		Tags:  tags,
	}

	// 空の場合は400を返す
	if data.Title == "" && data.Site == "" && len(data.Tags) == 0 {
		c.Status(400)
		c.SendString("Bad Request")
		return nil
	}

	// 画像生成
	img, err := lib.Generate(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		c.Status(500)
		c.SendString("Internal Server Error")
		return err
	}

	// jpegにエンコード
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		c.Status(500)
		c.SendString("Internal Server Error")
		return err
	}

	// レスポンスヘッダーの設定
	c.Set("Cache-Control", "public, max-age=86400")
	c.Set("Content-Type", "image/jpeg")
	c.Send(buffer.Bytes())
	return nil
}
