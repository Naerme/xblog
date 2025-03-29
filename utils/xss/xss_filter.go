package xss

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func XSSFilter(content string) (newcontent string) {
	// 文章正文防xss注入
	contentDoc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(content)))
	if err != nil {
		fmt.Println("防xss注入失败")
		return
	}
	contentDoc.Find("script").Remove()
	contentDoc.Find("img").Remove()
	contentDoc.Find("iframe").Remove()

	newcontent = contentDoc.Text()
	return
}
