package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/SaiSawant1/file-handler/filehandler"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	source      string
	destination string
	cursor      int
	options     []string
	textArea    textarea.Model
	viewport    viewport.Model
	selected    int
}

func Default() model {
	options := []string{"use default config", "custom src and destination"}
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()
	ta.Prompt = "┃ "
	ta.CharLimit = 280
	ta.SetWidth(30)
	ta.SetHeight(3)
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.ShowLineNumbers = false
	vp := viewport.New(30, 4)
	vp.SetContent("Enter your custom config")
	ta.KeyMap.InsertNewline.SetEnabled(false)
	return model{
		source:      "./",
		destination: "./",
		options:     options,
		textArea:    ta,
		viewport:    vp,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < 4 {
				m.cursor++
			}
		case "enter":
			m.selected = m.cursor
			if m.selected == 1 {
				m.textArea.View()
			}
			input := m.textArea.Value()
			if len(input) > 1 {
				m.viewport.SetContent(input)
				inputValue := strings.Split(input, " ")
				if len(inputValue) > 1 {
					filehandler.HandleFile(inputValue[0], inputValue[1])
				} else {
					filehandler.HandleFile(inputValue[0], "./")
				}

			}
		default:
			var cmd tea.Cmd
			m.textArea, cmd = m.textArea.Update(msg)
			return m, cmd
		}
	case cursor.BlinkMsg:
		var cmd tea.Cmd

		m.textArea, cmd = m.textArea.Update(msg)
		return m, cmd

	}

	return m, nil
}

func (m model) View() string {
	s := "Please Select an option\n"

	for i, choice := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if m.selected == i {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
		if m.selected == 1 {
			s = fmt.Sprintf("%s\n\n%s", m.viewport.View(), m.textArea.View())
		}
	}
	return s
}

func main() {
	p := tea.NewProgram(Default())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alias, there's been an error: %v", err)
		os.Exit(1)
	}
}
