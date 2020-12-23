package models

import "testing"

func TestGetAbsoluteURL(t *testing.T) {
	user := &User{
		Username: "damilola",
	}
	p := &Post{
		Slug:   "omo-e-be-like-123",
		Author: user,
	}
	testUrl := "/p/damilola/omo-e-be-like-123"
	url := p.GetAbsoluteURL()
	if url != testUrl {
		t.Errorf("Expected %s, but received %s", testUrl, url)
	}
}
