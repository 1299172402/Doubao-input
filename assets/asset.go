package assets

import (
	_ "embed"
)

//go:embed static/index.html
var IndexPage []byte

//go:embed static/logo.png
var LogoPNG []byte
