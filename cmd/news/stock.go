/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type news struct {
	Id        int    `json:"newsId"`
	Title     string `json:"title"`
	PublishAt int64  `json:"publishAt"`
}

type newsResponse struct {
	Items struct {
		LastPage int    `json:"last_page"`
		Data     []news `json:"data"`
	}
}

var newsMap = make(map[string][]news)

func getNews(stock string) {
	var b strings.Builder
	b.WriteString("https://api.cnyes.com/media/api/v1/newslist/TWS:")
	b.WriteString(stock)
	b.WriteString(":STOCK/symbolNews?page=1&limit=10")

	resp, err := http.Get(b.String())
	if err == nil {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			var newsResponse newsResponse
			err = json.Unmarshal(body, &newsResponse)
			if err == nil {
				newsMap[stock] = append(newsMap[stock], newsResponse.Items.Data...)
			}
		}
	}
	defer resp.Body.Close()
}

func printNews(idx int, publishAt int64, title string, newsId int) {
	fmt.Printf("%d. %s %s https://news.cnyes.com/news/id/%d\n", idx+1, time.Unix(publishAt, 0).Format("01/02"), title, newsId)
}

// stockCmd represents the stock command
var stockCmd = &cobra.Command{
	Use:   "stock [stock codes]",
	Short: "依照台股代號清單取得新聞",
	Long: `列出台股的相關新聞標題及其連結，台股代號順序不會依照輸入的順序。
台股代號清單請以逗號分隔，例如：2330,2317。`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stockArg := args[0]
		stockList := strings.Split(stockArg, ",")
		for _, stock := range stockList {
			getNews(stock)
		}
		// fmt.Printf("%+v\n", newsMap)
		for stock, newsList := range newsMap {
			fmt.Printf("------%s------\n", stock)
			for i, news := range newsList {
				printNews(i, news.PublishAt, news.Title, news.Id)
			}
		}
	},
}
