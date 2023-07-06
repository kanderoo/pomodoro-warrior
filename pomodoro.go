package main

// Pomodoro Timer Program

import tea "github.com/charmbracelet/bubbletea"

type model struct {
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return " "
}

func main() {
	p := tea.NewProgram(model{})
	p.Run()
}
