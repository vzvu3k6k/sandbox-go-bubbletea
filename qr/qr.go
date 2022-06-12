package main

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textInput textinput.Model
}

func newModel() model {
	ti := textinput.New()
	ti.Focus()

	return model{
		textInput: ti,
	}
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
		case "ctrl+d":
			if m.textInput.Value() == "" {
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m model) View() string {
	var out strings.Builder

	out.WriteString(m.textInput.View() + "\n")

	qr, err := renderQR(m.textInput.Value())
	if err != nil {
		out.WriteString(err.Error())
	} else {
		out.WriteString(qr)
	}

	return out.String()
}

var _ tea.Model = (*model)(nil)

func main() {
	m := newModel()
	prog := tea.NewProgram(m)
	if err := prog.Start(); err != nil {
		log.Fatal(err)
	}
}
