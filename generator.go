package main

import (
    "fmt"
    "github.com/russross/blackfriday/v2"
    "os"
    "io/ioutil"
    "html/template"
    "log"
)

type Post struct {
    Metadata PostMetadata
    Body template.HTML
}

type Index struct {
    PostsMetadata []PostMetadata
}

type PostMetadata struct {
    Title string
    Date string
    Link string
}
const raw_posts_directory_string = "./raw/"
const index_name = "blog/index.html"

func GenerateBlog() {
    files, _:= ioutil.ReadDir(raw_posts_directory_string)
    index := &Index{PostsMetadata: []PostMetadata{}}

    for _, file := range files {
        frontMatter, content := parse(raw_posts_directory_string + file.Name())

        body := blackfriday.Run(content)
        blogPostName := FormatBlogPostName(frontMatter["title"])
        link := CreateLink(blogPostName)
        fmt.Println(link)
        metadata := PostMetadata{Title: frontMatter["title"],
        Date: frontMatter["date"], Link: link}
        post := &Post{Metadata: metadata, Body: template.HTML(body)}

        t := template.Must(template.ParseFiles("templates/post.html"))

        f, err := os.Create(blogPostName + ".html")
        if err != nil {
            log.Println("create file: ", err)
            return
        }
        err = t.Execute(f, post)

        fmt.Println(file.Name())
        index.PostsMetadata = append(index.PostsMetadata, metadata)

        t = template.Must(template.ParseFiles("templates/index.html"))
        f, _ = os.Create(index_name)
        err = t.Execute(f, index)
        f.Close()
    }
}

