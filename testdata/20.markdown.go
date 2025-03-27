package main

import (
	"blogx_server/utils/markdowm"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var md = `

# 这是一级标题

> 这是应用

![xxx](jjdfnghjdf)

`

func main() {
	html := markdowm.MdToHtml(md)

	//fmt.Println(html)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
	if err != nil {
		fmt.Println(err)
		return
	}
	htmlText := doc.Text()
	fmt.Println(htmlText)
}
