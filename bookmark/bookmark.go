package bookmark

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Bookmark struct {
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
}

type BookmarkManager struct {
	bookmarks []Bookmark
	filePath  string
}

// NewBookmarkManager creates a new bookmark manager instance
func NewBookmarkManager() (*BookmarkManager, error) {
	// store the file in the user home dir (C:\Users\Rachit)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}

	filePath := filepath.Join(homeDir, ".bookmarks.json")

	bm := &BookmarkManager{
		bookmarks: make([]Bookmark, 0),
		filePath:  filePath,
	}

	// load the existing bookmarks from the JSON file
	err = bm.Load()
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to load bookmarks: %w", err)
		}
	}

	return bm, nil
}

// Add creates a new bookmark
func (bm *BookmarkManager) Add(name, path string) error {
	// use the current directory path if the path is not provided
	if path == "" {
		var err error
		path, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to fetch users current directory: %w", err)
		}
	}

	// convert the absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to resolve path: %w", err)
	}

	for _, b := range bm.bookmarks {
		if b.Name == name {
			// TODO ask the user whether to overwrite the existing path or not
			return fmt.Errorf("Path with the name '%s' already exists. \n", name)
		}
	}

	// create new bookmark
	bookmark := Bookmark{
		Name:      name,
		Path:      absPath,
		CreatedAt: time.Now(),
	}

	// add the bookmark to the saved bookmarks
	bm.bookmarks = append(bm.bookmarks, bookmark)

	return bm.Save()
}

// List returns all the bookmarks
func (bm *BookmarkManager) List() []Bookmark {
	return bm.bookmarks
}
