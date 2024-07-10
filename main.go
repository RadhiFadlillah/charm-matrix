package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	root := Root{
		Matrix: Matrix{
			ID:    "matrix-root",
			Speed: 50 * time.Millisecond,
		},
	}

	if _, err := tea.NewProgram(root, tea.WithAltScreen(), tea.WithMouseCellMotion()).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
