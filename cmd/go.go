package cmd

import (
	"fmt"
	"os"
	"rachitnimje/cli-bookmark-manager/bookmark"

	"github.com/spf13/cobra"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Navigate to a bookmarked directory",
	Long: `Navigate to a bookmarked directory. For example:

Examples:
  bm go bookmark_name     # Navigate to 'bookmark_name' bookmark
  bm go proj         	  # Navigate using fuzzy search (if multiple matches, shows list)`,
	Args: cobra.ExactArgs(1),
	Run:  runGo,
}

func init() {
	rootCmd.AddCommand(goCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runGo(cmd *cobra.Command, args []string) {
	// create a new BookmarkManager
	bm, err := bookmark.NewBookmarkManager()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	bookmarkName := args[0]

	// try for exact match of the query
	b, err := bm.Get(bookmarkName)
	if err != nil {
		fmt.Printf("Error: bookmark '%s' not found\n", bookmarkName)
		fmt.Printf("\nAvailable bookmarks\n")
		listBookmarks(bm)
		os.Exit(1)
	}
	fmt.Print(b.Path)

	// TODO: add fuzzy search
}

func listBookmarks(bm *bookmark.BookmarkManager) {
	bookmarks := bm.List()

	if len(bookmarks) == 0 {
		fmt.Println("No bookmarks found.")
		return
	}

	for _, b := range bookmarks {
		fmt.Printf("%s -> %s\n", b.Name, b.Path)
	}
}
