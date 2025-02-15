/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"github.com/spf13/cobra"
)

// NewsCmd represents the news command
var NewsCmd = &cobra.Command{
	Use:   "news",
	Short: "取得新聞，請使用子指令取得特定新聞",
}

func init() {
	NewsCmd.AddCommand(stockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// NewsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// NewsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
