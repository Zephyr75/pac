package game

import (
  "pac/ghost"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Point struct {
  X int
  Y int
}

type Game struct {
  Width int
  Height int
  PlayerPos Point
  PlayerChar string
  GameMap [][]int
  Ghosts []ghost.Ghost
}


func (g Game) Init() tea.Cmd {
  // g.ghosts = append(g.ghosts, ghost{point{g.Width - 2, g.Height - 2}, "B", blinky})
  // g.ghosts = append(g.ghosts, ghost{point{g.Width - 2, 1}, "I", inky})
  // g.ghosts = append(g.ghosts, ghost{point{1, g.Height - 2}, "P", pinky})
  // g.ghosts = append(g.ghosts, ghost{point{1, 1}, "Y", clyde})
  return nil
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "up", "k":
      if g.PlayerPos.Y > 0 && g.GameMap[g.PlayerPos.Y - 1][g.PlayerPos.X] != 2 {
        g.PlayerPos.Y--
      }
    case "down", "j":
      if g.PlayerPos.Y < g.Height - 1 && g.GameMap[g.PlayerPos.Y + 1][g.PlayerPos.X] != 2 {
        g.PlayerPos.Y++
      }
    case "left", "h":
      if g.PlayerPos.X > 0 && g.GameMap[g.PlayerPos.Y][g.PlayerPos.X - 1] != 2 {
        g.PlayerPos.X--
      }
    case "right", "l":
      if g.PlayerPos.X < g.Width - 1 && g.GameMap[g.PlayerPos.Y][g.PlayerPos.X + 1] != 2 {
        g.PlayerPos.X++
      }
    case "ctrl+c":
      return g, tea.Quit
    }
  }
  g.GameMap[g.PlayerPos.Y][g.PlayerPos.Y] = 0
  if g.PlayerChar == "C" { 
    g.PlayerChar = "c"
  } else {
    g.PlayerChar = "C"
  }
  return g, nil
}




var (
  yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffeb3b"))
  white = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
)

func (g Game) View() string {
  var s string
  for y := 0; y < g.Height; y++ {
    for x := 0; x < g.Width; x++ {
      if x == g.PlayerPos.X && y == g.PlayerPos.Y {
        s += yellow.Render(g.PlayerChar)
      } else {
        switch g.GameMap[y][x] {
        case 0:
          s += white.Render(" ")
        case 1:
          s += white.Render(".")
        case 2:
          s += white.Render("#")
        }
      }
    }
    s += "\n"
  }
  return s
}
