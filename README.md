# CLI Bookmark Manager

A lightweight command-line tool for managing directory bookmarks with fast navigation capabilities.

## Features

- 🔖 Bookmark frequently accessed directories
- ⚡️ Navigate instantly to bookmarked paths
- 📝 List all saved bookmarks
- 🗑️ Delete unwanted bookmarks
- 💾 Persistent storage across sessions
- 🔍 Fuzzy search support (planned)

## Installation

### Prerequisites
- Go 1.16+
- PowerShell 5.1+ (for directory navigation)

### Build and Install
```bash
# Clone the repository
git clone https://github.com/yourusername/cli-bookmark-manager.git
cd cli-bookmark-manager

# Build the executable
go build -o bm.exe 

# Install to your PATH (Windows example)
mkdir -p ~/bin
mv bm.exe ~/bin/

# Add ~/bin to PATH (if not already)
[Environment]::SetEnvironmentVariable("PATH", "$env:PATH;$HOME/bin", "User")