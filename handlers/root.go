package handlers

import (
	"fmt"
	"go-ogimg/lib"
	_ "image/png"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	// クエリ文字列のタグ部分のパース
	tags := []string{}
	if c.Request.FormValue("tags") != "" {
		tags = strings.Split(c.Request.FormValue("tags"), ",")
	}

	data := lib.SiteData{
		Title: c.Request.FormValue("title"),
		Site:  c.Request.FormValue("site"),
		Tags:  tags,
	}

	// 画像生成
	img, err := lib.Generate(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		c.Status(500)
		c.Writer.WriteString("Internal Server Error")
		return
	}

	// jpegに変換
	if err := lib.Png2jpg(img); err != nil {
		fmt.Fprintln(os.Stderr, err)
		c.Status(500)
		c.Writer.WriteString("Internal Server Error")
		return
	}

	// キャッシュ設定&レスポンス
	c.Header("cache-control", "public, max-age=86400")
	c.Writer.Write(img.Bytes())
}
