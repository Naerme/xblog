package flags

import (
	"blogx_server/models"
	"blogx_server/service/es_service"
)

func EsIndex() {
	article := models.ArticleModel{}
	es_service.CreateIndexV2(article.Index(), article.Mapping())
}
