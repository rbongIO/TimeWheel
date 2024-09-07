package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "支持单词格式转换",
	Long:  "支持各种单词格式转换",
	Run:   wordDealer,
}

func wordDealer(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}

func init() {}
