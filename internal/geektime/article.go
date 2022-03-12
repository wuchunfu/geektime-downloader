package geektime

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

type ArticleListResponse struct {
	Code int `json:"code"`
	Data struct {
		List []struct {
			ID               int    `json:"id"`
			ArticleTitle     string `json:"article_title"`
		} `json:"list"`
	} `json:"data"`
}

type ArticleSummary struct {
	AID   int
	Title string
}

func GetArticles(cid string, client *resty.Client) ([]ArticleSummary, error) {
	result := ArticleListResponse{}
	_, err := client.R().
		SetBody(
			map[string]interface{}{
				"cid":    cid,
				"order":  "earliest",
				"prev":   0,
				"sample": false,
				"size":   500,
			}).
		SetResult(&result).
		Post("/serv/v1/column/articles")

	if err != nil {
		return nil, err
	}

	if result.Code == 0 {
		var articles []ArticleSummary
		for _, v := range result.Data.List {
			articles = append(articles, ArticleSummary{
				AID:   v.ID,
				Title: v.ArticleTitle,
			})
		}
		return articles, nil
	} else {
		return nil, errors.New("get response failed")
	}
}
