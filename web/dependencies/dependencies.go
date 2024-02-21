package dependencies

import _ "embed"

var (
	//go:embed output.css
	OutputCss string

	//go:embed node_modules/htmx.org/dist/htmx.min.js
	HtmxJs string
)
