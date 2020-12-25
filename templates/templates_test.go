package templates

import (
	"testing"
)

func TestLineBreaks(t *testing.T) {
	content := "I want to repeat this at least one hundred times haha\r\nI want to repeat this at least one hundred times haha\r\n\r\nI want to repeat this at least one hundred times haha"
	html := LineBreaks(content)
	newContent := "<p>I want to repeat this at least one hundred times haha<br>I want to repeat this at least one hundred times haha</p><p>I want to repeat this at least one hundred times haha</p>"
	if string(html) != newContent {
		t.Errorf("Expected %s, but got %s", string(html), newContent)
	}
}

func Testmark(t *testing.T) {
	content := "I want to repeat this at least one hundred times"
	search := "want"
	html := mark(content, search)
	newContent := "I <mark>want</mark> to repeat this at least one hundred times"
	if string(html) != newContent {
		t.Errorf("Expected %s, but got %s", string(html), newContent)
	}
}

func TestTruncate(t *testing.T) {
	content := "I want to repeat this at least one hundred times"
	limit := 5
	content = Truncate(limit, content)
	newContent := "I wan..."
	if content != newContent {
		t.Errorf("Expected %s, but got %s", content, newContent)
	}
	if len(content) > len(newContent) {
		t.Errorf("Expected content length of at least %d, but got %d",
			len(newContent), len(content))
	}
}
