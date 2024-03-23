package main

import (
	"fmt"
	"log"
)

type Context struct {
	url     string
	content string
	data    any
}

type Handler func(*Context) error

func CheckURLHandler(c *Context) error {
	fmt.Printf("checking url: %s \n", c.url)
	return nil
}

func FetchContentHandler(c *Context) error {
	fmt.Printf("Fetching content from url: %s \n", c.url)
	c.content = "Hello World!!!"
	return nil
}

func ExtractDataHandler(c *Context) error {
	fmt.Printf("Extracting data from content: %s \n", c.content)
	c.data = map[string]interface{}{"title": "apple", "price": 10}
	return nil
}

func SaveDataHandler(c *Context) error {
	fmt.Printf("Saving data to database: %s \n", c.data)
	return nil
}

type HandlerNode struct {
	handler Handler
	next    *HandlerNode
}

func (node *HandlerNode) Handle(url string) error {
	ctx := Context{url: url}
	if node == nil || node.handler == nil {
		return nil
	}
	if err := node.handler(&ctx); err != nil {
		return err
	}
	return node.next.Handle(url)
}

func NewCrawler(handlers ...Handler) HandlerNode {
	node := HandlerNode{}
	if len(handlers) > 0 {
		node.handler = handlers[0]
	}

	currentNode := &node

	for i := 1; i < len(handlers); i++ {
		currentNode.next = &HandlerNode{handler: handlers[i]}
		currentNode = currentNode.next
	}

	return node
}

type WebCrawler struct {
	handler HandlerNode
}

func (wc WebCrawler) Crawler(url string) {
	if err := wc.handler.Handle(url); err != nil {
		log.Println(err)
	}
}

func main() {
	WebCrawler{
		handler: NewCrawler(
			CheckURLHandler,
			FetchContentHandler,
			ExtractDataHandler,
			SaveDataHandler,
		),
	}.Crawler("https://some-rich-content-website/some-page")
}
