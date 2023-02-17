package main

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed static/*
var static embed.FS

type tDat struct {
	Title string
	Site  string
	Tags  []string
}

func execTemplate(tDat tDat) (string, error) {
	t, err := template.ParseFS(static, "static/template.html")
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	if err := t.Execute(&buffer, tDat); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
