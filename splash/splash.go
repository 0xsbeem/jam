package splash

import (
	"os"
	"strings"

	"github.com/0xsbeem/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

var (
  docStyle = lipgloss.NewStyle()
  subtle = lipgloss.AdaptiveColor{Light: "#D9DCFF", Dark: "#383838"}
)

type Model struct {

}

func InitialModel() Model {
  return Model{}
}

func (m Model) Init() tea.Cmd {
  return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q":
      return m, tea.Quit
    }
  }
  return m, cmd
}

func (m Model) View() string {
  width, height, _ := term.GetSize(int(os.Stdout.Fd()))
  doc := strings.Builder{}
  banner := "     ██  █████  ███    ███\n     ██ ██   ██ ████  ████\n     ██ ███████ ██ ████ ██\n██   ██ ██   ██ ██  ██  ██\n █████  ██   ██ ██      ██\n\n"
  text := lipgloss.NewStyle().Width(26).Align(lipgloss.Center).Render(banner)


  ui := lipgloss.JoinVertical(lipgloss.Center, text)

  backgroundDot := docStyle.Copy().Foreground(subtle).Render(".")

  screen := lipgloss.Place(width-26, height,
    lipgloss.Center, lipgloss.Center,
    strings.ReplaceAll(ui, " ", backgroundDot),
    lipgloss.WithWhitespaceChars("."),
    lipgloss.WithWhitespaceForeground(subtle),
  )

  doc.WriteString(screen + "\n\n")

  if width > 0 {
    docStyle = docStyle.MaxWidth(width)
  }

  return docStyle.Render(doc.String())
  //return "Welcome to Jam"
}
