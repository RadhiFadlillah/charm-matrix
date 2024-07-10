package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Root struct {
	Matrix Matrix
}

func (r Root) Init() tea.Cmd {
	return tea.Batch(
		r.Matrix.Init(),
	)
}

func (r Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle root message
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC || msg.String() == "q" {
			return r, tea.Quit
		}

	case tea.WindowSizeMsg:
		return r, func() tea.Msg {
			return MatrixResized{
				ID:     r.Matrix.ID,
				Width:  msg.Width,
				Height: msg.Height,
			}
		}
	}

	// Handle children messages
	var cmd tea.Cmd
	var cmds tea.BatchMsg

	matrixModel, cmd := r.Matrix.Update(msg)
	cmds = append(cmds, cmd)

	// Assert child components
	r.Matrix = matrixModel.(Matrix)

	return r, tea.Batch(cmds...)
}

func (r Root) View() string {
	return r.Matrix.View()
}
