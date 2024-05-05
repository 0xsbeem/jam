package main

import (
	"fmt"
	"os"
	"time"

	splash "github.com/0xsbeem/jam/splash"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	// "strings"
	//
	// "github.com/0xsbeem/lipgloss"
	// "github.com/charmbracelet/bubbles/textinput"
	// oglg "github.com/charmbracelet/lipgloss"
	//  zone "github.com/lrstanley/bubblezone"
)

type model struct {
  showSplash bool
  splash splash.Model
  launchTimer timer.Model
}

func initialModel() model {
  showSplash := true
  splash := splash.InitialModel()
  launchTimer := timer.NewWithInterval(2, time.Second)

  return model{showSplash, splash, launchTimer}
}

func (m model) Init() tea.Cmd {
  return m.launchTimer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  switch msg := msg.(type) {
  case timer.TickMsg:
    m.launchTimer, cmd = m.launchTimer.Update(msg)
    return m, cmd

  case timer.TimeoutMsg:
    m.showSplash = false

  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q":
      return m, tea.Quit
    }
  }
  return m, cmd
}

func (m model) View() string {
  if m.showSplash {
    return m.splash.View()
  } else {
    return "This is not a splash."
  }
}

func main() {
  p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error: %v", err)
    os.Exit(1)
  }
}
