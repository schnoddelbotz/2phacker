package outputs

import "embed"

// content holds our static web server content.
//go:embed templates/*
var content embed.FS

func foo() {
	data, _ := content.ReadFile("index.html")
	print(string(data))
}
