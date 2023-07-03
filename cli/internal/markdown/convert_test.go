package markdown

import (
	"bytes"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte("Paragraph"), []byte("<p>Paragraph</p>\n")},
		{[]byte("# Heading"), []byte("<h1>Heading</h1>\n")},
		{[]byte("[Site](site.com)"), []byte("<p><a href=\"site.com\">Site</a></p>\n")},
	}
	for _, test := range tests {
		got, err := Convert(test.input)
		if err != nil {
			t.Errorf("Convert(%q) returned error: %v", test.input, err)
		}
		if !bytes.Equal(got, test.want) {
			t.Errorf("Convert(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
