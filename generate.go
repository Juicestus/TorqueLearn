package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	DIRECTORY = 0
	FILE      = 1
)

type NodeType int

type Node struct {
	Path     string
	Type     NodeType
	Children []Node
	Content  string
}

func NewNode(path string) Node {
	return Node{Path: path, Type: FILE, Children: []Node{}, Content: ""}
}

func GenerateStaticFiles() {
	node := ForEntryInDirectory("Juicestus/TorqueLearn/tree/master/pages/")
	node.PrintRecursive()
}

func ForEntryInDirectory(path string) Node {
	node := NewNode(path)

	res, err := http.Get("https://github.com/" + path)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	entries := doc.Find(".js-navigation-open")

	node.Type = FILE
	for i := range entries.Nodes {
		if i > 4 {
			href, exists := entries.Eq(i).Attr("href")
			if !exists {
				continue
			}

			node.Children = append(node.Children, ForEntryInDirectory(href))
			node.Type = DIRECTORY
		} else {
		}
	}

	if node.Type == FILE {
		// href, exists := doc.Find("#raw-url").Attr("href")
		// if !exists {

		// }
		// node.Content = ExtractContent(href)
		node.Content = ExtractContent(path)
	}
	return node
}

func ExtractContent(path string) string {
	path = strings.Replace(path, "/blob/", "/", 1)
	res, err := http.Get("https://raw.githubusercontent.com/" + path)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	fmt.Println("https://raw.githubusercontent.com" + path)

	if res.StatusCode != 200 {
		log.Fatalf("Invalid status code: %d %s", res.StatusCode, res.Status)
	}

	

	var buffer []byte
	_, err = res.Body.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buffer))

	return string(buffer)
}

func (n *Node) PrintRecursive() {
	if n.Type == DIRECTORY {
		log.Println(n.Path)
		for _, c := range n.Children {
			c.PrintRecursive()
		}
	}
}
