/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type categoryCommand struct {
	newsCommand
}

func (c *categoryCommand) getUrl(category string) string {
	var b strings.Builder
	b.WriteString("https://api.cnyes.com/media/api/v1/newslist/category/")
	b.WriteString(category)
	b.WriteString("?page=1&limit=10")
	return b.String()
}

func newCategoryCommand() *categoryCommand {
	return &categoryCommand{
		newsCommand: newsCommand{
			newsMap: make(map[string][]news),
		},
	}
}

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category [category codes]",
	Short: "依照新聞分類清單取得新聞",
	Long: `列出指定新聞分類的新聞標題及其連結，新聞分類順序不會依照輸入的順序。
新聞分類清單請以逗號分隔，例如：tw_stock,us_stock。`,
	Run: func(cmd *cobra.Command, args []string) {
		c := newCategoryCommand()
		categoryArg := args[0]
		categoryList := strings.Split(categoryArg, ",")
		for _, category := range categoryList {
			c.getNews(category, c.getUrl(category))
		}
		// fmt.Printf("%+v\n", newsMap)
		for category, newsList := range c.newsMap {
			fmt.Printf("------%s------\n", category)
			for i, news := range newsList {
				c.printNews(i, news.PublishAt, news.Title, news.Id)
			}
		}
	},
}
