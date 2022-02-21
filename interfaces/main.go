package main

import "fmt"

type Publisher interface {
	publish() error
}

type blogPost struct {
	postId int
	author string
	title  string
}

func (b blogPost) publish() error {
	fmt.Printf("The title on %s has been published by %s\n", b.title, b.author)
	return nil
}

func publishPost(pub Publisher) error {
	return pub.publish()
}

func main() {
	var pub Publisher
	pub = blogPost{1, "Diego", "Interfaces with golang!"}

	fmt.Println(pub)
	publishPost(pub)
}
