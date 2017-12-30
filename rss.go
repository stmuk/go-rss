package main

import (
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

func main() {
	file, _ := os.Open("rss")
	defer file.Close()
	fp := gofeed.NewParser()
	feed, _ := fp.Parse(file)
	fmt.Printf("%v", feed)

	for _, item := range feed.Items {

		fmt.Printf("%s %s\n", item.Title, item.GUID)

	}
}
