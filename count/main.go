package main

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	count int
	timer timer.Model
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		m.count = m.count + 1
		return m, cmd

	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case " ":
			m.count = m.count + 1
			return m, nil
		}
	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("count: %v, %v", m.count, m.timer.Interval)
}

var _ tea.Model = (*model)(nil)

func main() {
	m := model{
		count: 0,
		timer: timer.NewWithInterval(5*time.Second, time.Second),
	}
	prog := tea.NewProgram(m)
	err := prog.Start()
	if err != nil {
		log.Fatal(err)
	}
}
