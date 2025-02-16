/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"cnyes-stock-news/helper"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
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
				newsResponse.Items.Data = slices.DeleteFunc(newsResponse.Items.Data, func(n news) bool {
					return !helper.Filter(n.Title, include, exclude, regex)
				})
				c.newsMap[code] = append(c.newsMap[code], newsResponse.Items.Data...)
			}
		}
	}
	defer resp.Body.Close()
}

func (c *newsCommand) printNews(idx int, publishAt int64, title string, newsId int) {
	fmt.Printf("%d. %s %s https://news.cnyes.com/news/id/%d\n", idx+1, time.Unix(publishAt, 0).Format("01/02"), title, newsId)
}

var include []string
var exclude []string
var regex string

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
	NewsCmd.PersistentFlags().StringSliceVarP(&include, "include", "i", []string{}, "新聞標題須包含特定字串，以 , 分隔")
	NewsCmd.PersistentFlags().StringSliceVarP(&exclude, "exclude", "e", []string{}, "新聞標題不可包含特定字串，以 , 分隔")
	NewsCmd.PersistentFlags().StringVarP(&regex, "regex", "r", "", "新聞標題須符合正規表達式")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// NewsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
