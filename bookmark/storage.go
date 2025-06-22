package bookmark

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load reads the bookmarks from the JSON file
func (bm *BookmarkManager) Load() error {
	data, err := os.ReadFile(bm.filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &bm.bookmarks)
	if err != nil {
		return fmt.Errorf("failed to unmarshal bookmarks: %w", err)
	}

	return nil
}

// Save writes bookmarks to the JSON file
func (bm *BookmarkManager) Save() error {
	data, err := json.MarshalIndent(bm.bookmarks, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal bookmarks: %w", err)
	}

	// 0644 permission - readable by users,readable by group/others
	err = os.WriteFile(bm.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write bookmarks file: %w", err)
	}

	return nil
}
