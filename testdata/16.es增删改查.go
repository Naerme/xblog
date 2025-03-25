package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/models"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func create() {
	var article = models.ArticleModel{
		Model: models.Model{
			ID: 1,
		},
		Title:   "枫枫知道",
		Content: "这是内容",
		UserID:  1,
		Status:  1,
	}
	indexResponse, err := global.ESClient.Index().Index(article.Index()).BodyJson(article).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", indexResponse)
}
func list() {
	limit := 2
	page := 1
	from := (page - 1) * limit

	query := elastic.NewBoolQuery()
	res, err := global.ESClient.Search(models.ArticleModel{}.Index()).Query(query).From(from).Size(limit).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	count := res.Hits.TotalHits.Value // 总数
	fmt.Println(count)
	for _, hit := range res.Hits.Hits {
		fmt.Println(string(hit.Source))
	}

}

func DocDelete() {

	deleteResponse, err := global.ESClient.Delete().
		Index(models.ArticleModel{}.Index()).Id("1kNxy5UBk_gzP2OQjYpI").Refresh("true").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(deleteResponse)
}

func update() {
	global.ESClient.Update().Index(models.ArticleModel{}.Index()).Refresh("true").
		Id("10OBy5UBk_gzP2OQ0Yog").
		Doc(map[string]any{
			"content": "胡君寅123",
		}).Do(context.Background())
}

func main() {
	flags.Parse() //参数解析
	//fmt.Println(flags.FlagOptions)
	global.Conifg = core.ReadConf()
	core.InitLogrus()

	global.ESClient = core.EsConnect()
	//create()
	//list()
	//DocDelete()
	update()
}
