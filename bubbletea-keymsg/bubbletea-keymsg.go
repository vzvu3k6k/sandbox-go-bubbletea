package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct{}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		fmt.Printf("%s\n%#v\n", msg, msg)

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	return ""
}

func main() {
	p := tea.NewProgram(Model{}, tea.WithoutRenderer())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
