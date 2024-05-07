package static

import "embed"

//go:embed *
var EmbeddedStaticFS embed.FS
