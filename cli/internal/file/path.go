package file

import (
	"path/filepath"
	"strings"
)

const (
	MarkdownExt = ".md"
	HTMLExt     = ".html"
)

// IsMarkdown checks if a file path has the ".md" extension.
func IsMarkdown(path string) bool {
	return hasExt(path, MarkdownExt)

}

func hasExt(path, ext string) bool {
	return strings.EqualFold(filepath.Ext(path), ext)

}

// SetExt replaces the extension of a file path.
func SetExt(path, ext string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + ext
}
