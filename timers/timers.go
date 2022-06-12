package main

import (
	"log"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	spinner   spinner.Model
	textinput textinput.Model
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	cmds := make([]tea.Cmd, 0)

	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return m.spinner.View() + "\n" + m.textinput.View()
}

var _ tea.Model = (*model)(nil)

func main() {
	m := model{
		spinner:   spinner.New(),
		textinput: textinput.New(),
	}
	m.textinput.BlinkSpeed = time.Millisecond * 100
	m.textinput.Focus()

	prog := tea.NewProgram(m)
	if err := prog.Start(); err != nil {
		log.Fatal(err)
	}
}
