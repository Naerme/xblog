package markdown

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/sirupsen/logrus"
)

func MdToHtml(md string) string {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}

func ExtractContent(content string, len int) (abs string, err error) {
	// 把markdown转成html，再取文本
	html1 := MdToHtml(content)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(html1)))
	if err != nil {
		logrus.Errorf("提取正文失败：%s", err)
		return
	}
	htmlText := doc.Text()
	abs = htmlText
	if len > 200 {
		// 如果大于200，就取前200
		abs = string([]rune(htmlText)[:200])
	}
	return
}
