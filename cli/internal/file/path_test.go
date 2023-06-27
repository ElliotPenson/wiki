package file

import "testing"

func TestIsMarkdown(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{"markdown.md", true},
		{"relative/dir/markdown.md", true},
		{"/absolute/dir/markdown.md", true},
		{"other.txt", false},
		{"relative/dir/other.txt", false},
		{"/absolute/dir/other.txt", false},
		{"relative/dir", false},
		{"/absolute/dir", false},
	}
	for _, test := range tests {
		if got := IsMarkdown(test.path); got != test.want {
			t.Errorf("IsMarkdown(%q) = %t, want %t", test.path, got, test.want)
		}
	}
}

func TestSetExt(t *testing.T) {
	tests := []struct {
		path string
		ext  string
		want string
	}{
		{"test.md", ".html", "test.html"},
		{"leading/dir/test.md", ".html", "leading/dir/test.html"},
		{"/home/dir/test.md", ".html", "/home/dir/test.html"},
	}
	for _, test := range tests {
		if got := SetExt(test.path, test.ext); got != test.want {
			t.Errorf("SetExt(%q, %q) = %q, want %q", test.path, test.ext, got, test.want)
		}
	}
}
