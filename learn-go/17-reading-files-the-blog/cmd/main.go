package main

import (
	"log"
	"os"

	"github.com/lin-br/go-lin/blogposts/blogposts"
)

func main() {
	path, _ := os.Getwd()
	dir := os.DirFS(path + "/posts")
	posts, err := blogposts.NewPostsFromFS(dir)
	if err != nil {
		log.Fatal("failed to get posts:", err)
	}
	log.Println(posts)
}
