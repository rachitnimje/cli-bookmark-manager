package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"rachitnimje/cli-bookmark-manager/bookmark"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runList,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runList(cmd *cobra.Command, args []string) {
	// create a new BookmarkManager
	bm, err := bookmark.NewBookmarkManager()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	bookmarks := bm.List()

	for _, b := range bookmarks {
		fmt.Printf("Name: %v Path: %v \n", b.Name, b.Path)
	}
}
