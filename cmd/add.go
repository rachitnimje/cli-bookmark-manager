package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"rachitnimje/cli-bookmark-manager/bookmark"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [name] [path]",
	Short: "Adds the path as a new bookmark",
	Long: `Adds the path provided as a new bookmark to the cli-bookmark-manager. 
For example:
	bm add bookmark_name							# bookmarks current directory as "bookmark_name"
	bm add bookmark_name path/to/the/project		# bookmarks specified directory as "bookmark_name"`,
	Args: cobra.RangeArgs(1, 2),
	Run:  runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runAdd(cmd *cobra.Command, args []string) {
	// create a new BookmarkManager
	bm, err := bookmark.NewBookmarkManager()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// extract the name and the path from the args
	name := args[0]
	path := ""

	if len(args) > 1 {
		path = args[1]
	}

	// save the bookmark to the JSON file
	err = bm.Add(name, path)
	if err != nil {
		fmt.Printf("Error adding bookmark: %s", err)
		os.Exit(1)
	}

	if path == "" {
		currentDir, _ := os.Getwd()
		fmt.Printf("✓ Bookmarked current directory (%s) as '%s'\n", currentDir, name)
	} else {
		fmt.Printf("✓ Bookmarked '%s' as '%s'\n", path, name)
	}
}
