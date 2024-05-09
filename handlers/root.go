package handlers

import (
	"fmt"
	"go-ogimg/lib"
	_ "image/png"
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

	// 画像生成
	img, err := lib.Generate(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		c.Status(500)
		c.SendString("Internal Server Error")
		return err
	}

	// jpegに変換
	if err := lib.Png2jpg(img); err != nil {
		fmt.Fprintln(os.Stderr, err)
		c.Status(500)
		c.SendString("Internal Server Error")
		return err
	}

	c.Response().Header.Add("cache-control", "public, max-age=86400")
	c.Send(img.Bytes())
	return nil
}
