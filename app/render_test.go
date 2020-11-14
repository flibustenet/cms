package app

import (
	"strings"
	"testing"
)

func TestMarkdown(t *testing.T) {
	input := `**gras**`
	output := string(markdown(input))
	if !strings.Contains(output, "<p><strong>gras</strong></p>") {
		t.Fatalf("Markdown de %s != '%s'", input, output)
	}
}
