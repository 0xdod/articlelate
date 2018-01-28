package templates

import "testing"

func TestLineBreaks(t *testing.T) {
	content := "I want to repeat this at least one hundred times haha\r\nI want to repeat this at least one hundred times haha\r\n\r\nI want to repeat this at least one hundred times haha"
	html := LineBreaks(content)
	newContent := "<p>I want to repeat this at least one hundred times haha<br>I want to repeat this at least one hundred times haha</p><p>I want to repeat this at least one hundred times haha</p>"
	if string(html) != newContent {
		t.Errorf("Expected %s, but got %s", string(html), newContent)
	}
}
