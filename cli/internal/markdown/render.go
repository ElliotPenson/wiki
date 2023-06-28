package markdown

import (
	"bytes"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
)

var md = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		extension.Footnote,
		highlighting.Highlighting,
	),
)

// Render converts Markdown into HTML.
func Render(markdown []byte) ([]byte, error) {
	var b bytes.Buffer
	if err := md.Convert(markdown, &b); err != nil {
		return []byte{}, err
	}
	return b.Bytes(), nil
}
