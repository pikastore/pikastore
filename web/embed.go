package web

import (
	"embed"
)

//go:embed build/*
var WebAssets embed.FS

func GetWebAsset(path string) ([]byte, error) {
	return WebAssets.ReadFile("build" + path)
}

func ListWebAssets() ([]string, error) {
	entries, err := WebAssets.ReadDir("build")
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		files = append(files, entry.Name())
	}
	return files, nil
}
