package assets

import _ "embed"

var (
	//go:embed output.css
	OutputCss string

	//go:embed htmx.min.js
	HtmxJs string
)
