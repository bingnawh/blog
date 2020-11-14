package main

import "strings"

func FormatBlogPostName(blogPostName string) string {
	return strings.ReplaceAll(strings.TrimSpace(blogPostName), " ", "-")

}

func CreateLink(blogPostName string) string {
	blogPostName = FormatBlogPostName(blogPostName)
	link := blogPostName + ".html"
	return link
}
