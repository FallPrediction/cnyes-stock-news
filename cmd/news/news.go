/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		Data []news `json:"data"`
	}
}

type newsCommand struct {
	newsMap map[string][]news
}

func (c *newsCommand) getNews(code string, url string) {
	resp, err := http.Get(url)
	if err == nil {
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			var newsResponse newsResponse
			err = json.Unmarshal(body, &newsResponse)
			if err == nil {
				c.newsMap[code] = append(c.newsMap[code], newsResponse.Items.Data...)
			}
		}
	}
	defer resp.Body.Close()
}

func (c *newsCommand) printNews(idx int, publishAt int64, title string, newsId int) {
	fmt.Printf("%d. %s %s https://news.cnyes.com/news/id/%d\n", idx+1, time.Unix(publishAt, 0).Format("01/02"), title, newsId)
}

// NewsCmd represents the news command
var NewsCmd = &cobra.Command{
	Use:   "news",
	Short: "取得新聞，請使用子指令取得特定新聞",
}

func init() {
	NewsCmd.AddCommand(stockCmd)
	NewsCmd.AddCommand(categoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// NewsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// NewsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
