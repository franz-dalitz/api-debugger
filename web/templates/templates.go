package templates

import "embed"

var (
	//go:embed components/* layouts/* pages/*
	templates embed.FS
)

func Get() embed.FS {
	return templates
}
