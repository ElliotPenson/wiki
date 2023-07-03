package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/elliotpenson/wiki/cli/internal/file"
	"github.com/elliotpenson/wiki/cli/internal/markdown"
	"github.com/spf13/cobra"
)

const (
	filePerm = 0644
	dirPerm  = 0755
)

var (
	src  string
	dest string
)

var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Convert wiki markdown into HTML",
	RunE: func(cmd *cobra.Command, args []string) error {
		return render(src, dest)
	},
}

func init() {
	renderCmd.Flags().StringVarP(&src, "src", "s", "pages", "Directory with Markdown input")
	renderCmd.Flags().StringVarP(&dest, "dest", "d", "public", "Directory for HTML output")
	rootCmd.AddCommand(renderCmd)
}

func render(src, dest string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		to := filepath.Join(dest, rel)
		if d.IsDir() {
			return os.Mkdir(to, dirPerm)
		}
		return renderFile(path, to)
	})
}

func renderFile(src, dest string) error {
	if file.IsMarkdown(src) {
		htmlPath := file.SetExt(dest, file.HTMLExt)
		return renderMarkdown(src, htmlPath)
	}
	// Non-markdown files (e.g. images) are copied as-is.
	return file.Copy(src, dest)
}

func renderMarkdown(src, dest string) error {
	md, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("error reading Markdown file %q: %w", src, err)
	}
	html, err := markdown.Render(md)
	if err != nil {
		return fmt.Errorf("error rendering Markdown %q: %w", src, err)
	}
	return os.WriteFile(dest, html, filePerm)
}
