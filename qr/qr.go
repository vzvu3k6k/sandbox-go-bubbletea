package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	data string
}

func (model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "Bubble, bubble..."
}

var _ tea.Model = (*model)(nil)

func main() {
	m := model{}
	prog := tea.NewProgram(m)
	if err := prog.Start(); err != nil {
		log.Fatal(err)
	}
}
