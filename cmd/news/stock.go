/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type stockCommand struct {
	newsCommand
}

func (c *stockCommand) getUrl(stock string) string {
	var b strings.Builder
	b.WriteString("https://api.cnyes.com/media/api/v1/newslist/TWS:")
	b.WriteString(stock)
	b.WriteString(":STOCK/symbolNews?page=1&limit=10")
	return b.String()
}

func newStockCommand() *stockCommand {
	return &stockCommand{
		newsCommand: newsCommand{
			newsMap: make(map[string][]news),
		},
	}
}

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock [stock codes]",
	Short: "依照台股代號清單取得新聞",
	Long: `列出台股的相關新聞標題及其連結，台股代號順序不會依照輸入的順序。
台股代號清單請以逗號分隔，例如：2330,2317。`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := newStockCommand()
		stockArg := args[0]
		stockList := strings.Split(stockArg, ",")
		for _, stock := range stockList {
			c.getNews(stock, c.getUrl(stock))
		}
		// fmt.Printf("%+v\n", newsMap)
		for stock, newsList := range c.newsMap {
			fmt.Printf("------%s------\n", stock)
			for i, news := range newsList {
				c.printNews(i, news.PublishAt, news.Title, news.Id)
			}
		}
	},
}
