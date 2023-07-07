package main

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

const pomodoroLength = time.Second * 2
const shortBreakLength = time.Second * 1
const longBreakLength = time.Second * 1
const longBreakInterval = 4

const (
	Pomodoro   = iota // 0
	ShortBreak = iota // 1
	LongBreak  = iota // 2
)

// Struct Definitions

type model struct {
	mode          int
	pomodoroCount int
	timer         timer.Model
	keymap        keymap
}

type keymap struct {
	timerToggle key.Binding
	reset       key.Binding
	quit        key.Binding
}

// Elm Architecture Methods

func (m model) Init() tea.Cmd {
	// stop the timer from starting upon start, wait for the user
	return m.timer.Stop()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case timer.StartStopMsg, timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.TimeoutMsg:
		switch m.mode {
		case Pomodoro:
			m.pomodoroCount++
			if m.pomodoroCount%longBreakInterval == 0 {
				m.mode = LongBreak
				m.timer = timer.New(longBreakLength)
			} else {
				m.mode = ShortBreak
				m.timer = timer.New(shortBreakLength)
			}
		case ShortBreak, LongBreak:
			m.mode = Pomodoro
			m.timer = timer.New(pomodoroLength)
		}
		return m, m.timer.Stop()

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.timerToggle):
			// Toggle the timer
			return m, m.timer.Toggle()
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Pomodoro Timer\n\n"

	// Print mode and remaining time
	switch m.mode {
	case Pomodoro:
		s += "Pomodoro"
	case ShortBreak:
		s += "Short Break"
	case LongBreak:
		s += "Long Break"
	}
	s += ": " + m.timer.View() + "\n\n"

	// Print timer status
	s += "Timer is currently "
	if m.timer.Running() {
		s += "running"
	} else {
		s += "stopped"
	}

	return s
}

func main() {
	p := tea.NewProgram(model{
		mode:          Pomodoro,
		pomodoroCount: 0,
		timer:         timer.NewWithInterval(pomodoroLength, time.Second),
		keymap: keymap{
			timerToggle: key.NewBinding(key.WithKeys(" ")),
			reset:       key.NewBinding(key.WithKeys("r")),
			quit:        key.NewBinding(key.WithKeys("q", "ctrl+c")),
		},
	})
	p.Run()
}
