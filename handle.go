package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"strings"
)

func ogimageHandle(w http.ResponseWriter, r *http.Request) {
	tags := []string{}
	if r.FormValue("tags") != "" {
		tags = strings.Split(r.FormValue("tags"), ",")
	}

	data := siteData{
		Title: r.FormValue("title"),
		Site:  r.FormValue("site"),
		Tags:  tags,
	}

	img, err := generate(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		errorResponse(w)
		return
	}

	if err := png2jpg(img); err != nil {
		fmt.Fprintln(os.Stderr, err)
		errorResponse(w)
		return
	}

	w.Header().Add("cache-control", "public, max-age=86400")
	w.Write(img.Bytes())
}

func png2jpg(buf *bytes.Buffer) error {
	img, _, err := image.Decode(buf)
	if err != nil {
		return err
	}
	jpeg.Encode(buf, img, &jpeg.Options{Quality: 90})
	return nil
}
