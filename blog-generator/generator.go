package main

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type Post struct {
	Metadata PostMetadata
	Body     template.HTML
}

type Index struct {
	PostsMetadata []PostMetadata
}

type PostMetadata struct {
	Title string
	Date  string
	Link  string
}

const raw_posts_directory_string = "./raw/"
const blog_name = "blog.html"

func createIndex(index Index, indexFileName string) {
	fmt.Println("Creating blog.html...")
	t := template.Must(template.ParseFiles("templates/blog.html"))
	f, _ := os.Create(indexFileName)
	defer f.Close()
	_ = t.Execute(f, index)
}
func GenerateBlog() {
	files, _ := ioutil.ReadDir(raw_posts_directory_string)
	index := Index{PostsMetadata: []PostMetadata{}}

	for _, file := range files {
		frontMatter, content := parse(raw_posts_directory_string + file.Name())

		body := blackfriday.Run(content)
		blogPostName := FormatBlogPostName(frontMatter["title"])
		link := CreateLink(blogPostName)
		log.Println("Creating link for blog.html at ", link)
		metadata := PostMetadata{Title: frontMatter["title"],
			Date: frontMatter["date"], Link: link}
		post := &Post{Metadata: metadata, Body: template.HTML(body)}

		t := template.Must(template.ParseFiles("templates/post.html"))

		f, err := os.Create(blogPostName + ".html")
		defer f.Close()
		if err != nil {
			log.Println("create file: ", err)
			return
		}
		_ = t.Execute(f, post)

		index.PostsMetadata = append(index.PostsMetadata, metadata)

	}
	indexFileName := blog_name
	createIndex(index, indexFileName)
}
