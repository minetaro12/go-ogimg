package handlers

import (
	"fmt"
	"go-ogimg/lib"
	_ "image/png"
	"net/http"
	"os"
	"strings"
)

func Root(w http.ResponseWriter, r *http.Request) {
	// クエリ文字列のタグ部分のパース
	tags := []string{}
	if r.FormValue("tags") != "" {
		tags = strings.Split(r.FormValue("tags"), ",")
	}

	data := lib.SiteData{
		Title: r.FormValue("title"),
		Site:  r.FormValue("site"),
		Tags:  tags,
	}

	// 画像生成
	img, err := lib.Generate(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// jpegに変換
	if err := lib.Png2jpg(img); err != nil {
		fmt.Fprintln(os.Stderr, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// キャッシュ設定&レスポンス
	w.Header().Add("cache-control", "public, max-age=86400")
	w.Write(img.Bytes())
}
