/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

type stockCommand struct {
	newsCommand
}

func (c *stockCommand) getUrl(stock string, page int) string {
	var b strings.Builder
	b.WriteString("https://api.cnyes.com/media/api/v1/newslist/TWS:")
	b.WriteString(stock)
	b.WriteString(fmt.Sprintf(":STOCK/symbolNews?page=%d&limit=10", page))
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
	Long: `列出過去 24 小時發布的台股相關新聞標題及其連結，最多 30 篇，台股代號順序不會依照輸入的順序。
台股代號清單請以逗號分隔，例如：2330,2317。`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := newStockCommand()
		stockArg := args[0]
		stockList := strings.Split(stockArg, ",")
		past := time.Now().Add(time.Hour * -24).Unix()
		ch := make(chan struct{}, 2)
		var wg sync.WaitGroup
		wg.Add(len(stockList))
		for _, stock := range stockList {
			go func() {
				ch <- struct{}{}
				defer wg.Done()
				for page := 1; page <= 3; page++ {
					if moreNews := c.getNews(stock, c.getUrl(stock, page), past); !moreNews {
						break
					}
				}
				c.filterNews(stock, past)
				// Pause for 1 second between every two stocks to avoid too many requests.
				time.Sleep(1 * time.Second)
				<- ch
			}()
		}
		wg.Wait()
		// fmt.Printf("%+v\n", newsMap)
		for stock, newsList := range c.newsMap {
			fmt.Printf("------%s------\n", stock)
			for i, news := range newsList {
				c.printNews(i, news.PublishAt, news.Title, news.Id)
			}
		}
	},
}
