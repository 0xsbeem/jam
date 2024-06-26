package main

import (
	"fmt"
	"os"
	"time"

	splash "github.com/0xsbeem/jam/splash"
 	"github.com/0xsbeem/stickers"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	// "strings"
	//
	"github.com/0xsbeem/lipgloss"
	// "github.com/charmbracelet/bubbles/textinput"
	// oglg "github.com/charmbracelet/lipgloss"
	//  zone "github.com/lrstanley/bubblezone"
)

var (
	style1  = lipgloss.NewStyle().Background(lipgloss.Color("#fc5c65"))
	style2  = lipgloss.NewStyle().Background(lipgloss.Color("#fd9644"))
	style3  = lipgloss.NewStyle().Background(lipgloss.Color("#fed330"))
	style4  = lipgloss.NewStyle().Background(lipgloss.Color("#26de81"))
	style5  = lipgloss.NewStyle().Background(lipgloss.Color("#2bcbba"))
	style6  = lipgloss.NewStyle().Background(lipgloss.Color("#eb3b5a"))
	style7  = lipgloss.NewStyle().Background(lipgloss.Color("#fa8231"))
	style8  = lipgloss.NewStyle().Background(lipgloss.Color("#f7b731"))
	style9  = lipgloss.NewStyle().Background(lipgloss.Color("#20bf6b"))
	style10 = lipgloss.NewStyle().Background(lipgloss.Color("#0fb9b1"))
	style11 = lipgloss.NewStyle().Background(lipgloss.Color("#45aaf2"))
	style12 = lipgloss.NewStyle().Background(lipgloss.Color("#4b7bec"))
	style13 = lipgloss.NewStyle().Background(lipgloss.Color("#a55eea"))
	style14 = lipgloss.NewStyle().Background(lipgloss.Color("#d1d8e0"))
	style15 = lipgloss.NewStyle().Background(lipgloss.Color("#778ca3"))
	style16 = lipgloss.NewStyle().Background(lipgloss.Color("#2d98da"))
	style17 = lipgloss.NewStyle().Background(lipgloss.Color("#3867d6"))
	style18 = lipgloss.NewStyle().Background(lipgloss.Color("#8854d0"))
	style19 = lipgloss.NewStyle().Background(lipgloss.Color("#a5b1c2"))
	style20 = lipgloss.NewStyle().Background(lipgloss.Color("#4b6584"))
)

type model struct {
  showSplash bool
  splash splash.Model
  launchTimer timer.Model
 	flexBox *stickers.FlexBox
}

func initialModel() model {
  showSplash := true
  splash := splash.InitialModel()
  launchTimer := timer.NewWithInterval(2, time.Second)
  flex := stickers.NewFlexBox(0, 0)

  rows := []*stickers.FlexBoxRow{
    flex.NewRow().AddCells(
      []*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(1, 2).SetStyle(style1),
				stickers.NewFlexBoxCell(1, 3).SetStyle(style2),
				stickers.NewFlexBoxCell(1, 4).SetStyle(style3),
				stickers.NewFlexBoxCell(1, 3).SetStyle(style2),
				stickers.NewFlexBoxCell(1, 2).SetStyle(style1),
      },
    ),
    flex.NewRow().AddCells(
      []*stickers.FlexBoxCell{
        stickers.NewFlexBoxCell(1, 20).SetStyle(style20),
      },
    ),
  }

  flex.AddRows(rows)

  return model{showSplash, splash, launchTimer, flex}
}

func (m model) Init() tea.Cmd {
  return m.launchTimer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    m.flexBox.SetWidth(msg.Width)
    m.flexBox.SetHeight(msg.Height)
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
    return m.flexBox.Render()
  }
}

func main() {
  p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error: %v", err)
    os.Exit(1)
  }
}
