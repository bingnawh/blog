package main

import "testing"

func TestFormatBlogPostName(t *testing.T) {
	blogTitle := " sample title "
	got := FormatBlogPostName(blogTitle)
	want := "sample-title"
	if got != want {
		t.Errorf("got = %s, want %s", got, want)
	}
}
func TestCreateLink(t *testing.T) {
	formattedBlogTitle := "Sample-Title"
	got := CreateLink(formattedBlogTitle)
	want := "sample-title"
	if got != want {
		t.Errorf("got = %s, want %s", got, want)
	}
}
