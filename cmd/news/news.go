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

func (c *newsCommand) sendRequest(url string) newsResponse {
	var (
		retries      int = 3
		resp         *http.Response
		err          error
		newsResponse newsResponse
	)
	for retries > 0 {
		resp, err = http.Get(url)
		if err != nil || resp.StatusCode != 200 {
			retries--
			continue
		} else {
			break
		}
	}
	body, err := io.ReadAll(resp.Body)
	if err == nil {
		json.Unmarshal(body, &newsResponse)
	}
	defer resp.Body.Close()
	return newsResponse
}

// Send request to get news.
// The return value means there could be more news published after the parameter publishAt.
func (c *newsCommand) getNews(code string, url string, publishAt int64) bool {
	newsResponse := c.sendRequest(url)
	if len(newsResponse.Items.Data) == 0 {
		return false
	}
	c.newsMap[code] = append(c.newsMap[code], newsResponse.Items.Data...)
	return c.newsMap[code][len(c.newsMap[code])-1].PublishAt > publishAt
}

func (c *newsCommand) printNews(idx int, publishAt int64, title string, newsId int) {
	fmt.Printf("%d. %s %s https://news.cnyes.com/news/id/%d\n", idx+1, time.Unix(publishAt, 0).Format("01/02"), title, newsId)
}

func (c *newsCommand) filterNews(code string, publishAt int64) {
	c.newsMap[code] = slices.DeleteFunc(c.newsMap[code], func(n news) bool {
		// Remove news that is not matched the rules and has been posted for more than 24 hours
		return !helper.Filter(n.Title, include, exclude, regex) || n.PublishAt < publishAt
	})
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
