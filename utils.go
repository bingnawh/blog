package main

import "strings"
func FormatBlogPostName(blogPostName string) string {
    return strings.ReplaceAll(strings.TrimSpace(blogPostName), " ", "-")

}

func CreateLink(blogPostName string) string {
    const posts_dir = "blog/posts/"
    blogPostName = FormatBlogPostName(blogPostName)
    link := posts_dir + blogPostName + ".html"
    return link
}
