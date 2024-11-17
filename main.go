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

type ViewScreen int

const (
	HomeScreen ViewScreen = iota
	CustomScreen
	FileScreen
)

type model struct {
	source      string
	destination string
	cursor      int
	options     []string
	textArea    textarea.Model
	viewport    viewport.Model
	currScreen  ViewScreen
	msgChannel  chan string
	outPut      string
}

func Default() model {
	options := []string{"Use default Source and Destination ('./','./')", "Provide a Custom Source and Destination"}
	ta := textarea.New()
	ta.Placeholder = "Pick the sorce and destination..."
	ta.Focus()
	ta.Prompt = "┃ "
	ta.CharLimit = 280
	ta.SetWidth(40)
	ta.SetHeight(1)
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.ShowLineNumbers = false
	vp := viewport.New(100, 1)
	ta.KeyMap.InsertNewline.SetEnabled(false)
	msgCh := make(chan string)
	return model{
		source:      "./",
		destination: "./",
		options:     options,
		textArea:    ta,
		viewport:    vp,
		currScreen:  HomeScreen,
		msgChannel:  msgCh,
		outPut:      "",
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
			if m.cursor < 1 {
				m.cursor++
			}
		case "enter":
			switch m.cursor {
			case 0:
				m.currScreen = FileScreen
				go func() {
					filehandler.HandleFile(m.source, m.destination, m.msgChannel)
				}()
			case 1:
				m.currScreen = CustomScreen
				m.textArea.View()
				input := m.textArea.Value()
				if len(input) > 1 {
					m.viewport.SetContent(input)
					inputValue := strings.Split(input, " ")
					if len(inputValue) > 1 {
						m.source = inputValue[0]
						m.destination = inputValue[1]
						m.currScreen = FileScreen
						go func() {
							filehandler.HandleFile(m.source, m.destination, m.msgChannel)
						}()
					} else {
						m.source = inputValue[0]
						m.currScreen = FileScreen
						go func() {
							filehandler.HandleFile(m.source, m.destination, m.msgChannel)
						}()
					}
				}
			}
		default:
			var cmd tea.Cmd
			m.textArea, cmd = m.textArea.Update(msg)
			return m, cmd
		}
	case cursor.BlinkMsg:
		var cmd tea.Cmd
		select {
		case val, ok := <-m.msgChannel:
			if ok {
				m.outPut += fmt.Sprintf("%s\n", val)

			}
		default:
		}
		m.textArea, cmd = m.textArea.Update(msg)
		return m, cmd

	}

	return m, nil
}

func (m model) View() string {

	var s string

	switch m.currScreen {
	case HomeScreen:
		s += fmt.Sprintf("Please use arrow keys up and down or j and k to navigate:-\n\n")
		for i, option := range m.options {

			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			checked := " "
			if int(m.currScreen) == i {
				checked = "x"
			}
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, option)
		}
		s += fmt.Sprintf("\npress ctrl+c or q to quite\n")
	case CustomScreen:
		m.viewport.SetContent("Please enter space seperated source and destination ('./src ./dest') \n if destination folder is not found the destination folder will be created")
		s = fmt.Sprintf("%s\n\n%s", m.viewport.View(), m.textArea.View())
	case FileScreen:
		m.viewport.SetContent("work in progress..\n")
		s += m.outPut
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
