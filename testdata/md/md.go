package main

import (
	"blogx_server/models"
	"blogx_server/service/text_service"
	"fmt"
	"os"
)

func main() {
	byteData, _ := os.ReadFile("text.md")
	list := text_service.MdContentTransformation(models.ArticleModel{
		Model:   models.Model{ID: 1},
		Title:   "xxx",
		Content: string(byteData),
	})

	fmt.Println(list)

}
