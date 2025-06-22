/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"rachitnimje/cli-bookmark-manager/bookmark"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a bookmark by name",
	Long: `Delete a bookmark by its name. For example:
	bm delete my-project`,
	Args: cobra.ExactArgs(1),
	Run:  runDelete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runDelete(cmd *cobra.Command, args []string) {
	// create a new BookmarkManager
	bm, err := bookmark.NewBookmarkManager()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	name := args[0]
	err = bm.Delete(name)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Bookmark '%s' deleted successfully \n", name)
}
