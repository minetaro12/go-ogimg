package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func ogimageHandle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var tags []string
	if r.FormValue("tags") != "" {
		tags = strings.Split(r.FormValue("tags"), ",")
	}

	dat := tDat{
		Title: r.FormValue("title"),
		Site:  r.FormValue("site"),
		Tags:  tags,
	}

	html, err := execTemplate(dat)
	if err != nil {
		errorResponse(w)
		return
	}

	var imgBuf []byte
	if err := chromedp.Run(ctx, takeScreenshot(html, 90, &imgBuf)); err != nil {
		errorResponse(w)
		return
	}

	w.Header().Add("cache-control", "public, max-age=86400")
	w.Write(imgBuf)
}

func takeScreenshot(html string, quality int, buffer *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EmulateViewport(1200, 630),
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, html).Do(ctx)
		}),
		chromedp.WaitVisible("body > .loaded"),
		chromedp.FullScreenshot(buffer, quality),
	}
}
