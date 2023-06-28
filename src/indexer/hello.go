package main

import (
	"fmt"
	"indexer/util"
)

func main() {
	go util.GetCrawlPage()
	fmt.Println(util.GetCrawlPage())

}
