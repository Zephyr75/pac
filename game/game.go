package game

import (
  "pac/ghost"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Point struct {
  x int
  y int
}

type Game struct {
  width int
  height int
  playerPos Point
  playerChar string
  gameMap [][]int
  ghosts []ghost.Ghost
}


func (g Game) Init() tea.Cmd {
  // g.ghosts = append(g.ghosts, ghost{point{g.width - 2, g.height - 2}, "B", blinky})
  // g.ghosts = append(g.ghosts, ghost{point{g.width - 2, 1}, "I", inky})
  // g.ghosts = append(g.ghosts, ghost{point{1, g.height - 2}, "P", pinky})
  // g.ghosts = append(g.ghosts, ghost{point{1, 1}, "Y", clyde})
  return nil
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "up", "k":
      if g.playerPos.y > 0 {
        g.playerPos.y--
      }
    case "down", "j":
      if g.playerPos.y < g.height - 1 {
        g.playerPos.y++
      }
    case "left", "h":
      if g.playerPos.x > 0 {
        g.playerPos.x--
      }
    case "right", "l":
      if g.playerPos.x < g.width - 1 {
        g.playerPos.x++
      }
    case "ctrl+c":
      return g, tea.Quit
    }
  }
  g.gameMap[g.playerPos.y][g.playerPos.x] = 0
  if g.playerChar == "C" { 
    g.playerChar = "c"
  } else {
    g.playerChar = "C"
  }
  return g, nil
}




var (
  yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffeb3b"))
  white = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
)

func (g Game) View() string {
  var s string
  for y := 0; y < g.height; y++ {
    for x := 0; x < g.width; x++ {
      if x == g.playerPos.x && y == g.playerPos.y {
        s += yellow.Render(g.playerChar)
      } else {
        switch g.gameMap[y][x] {
        case 0:
          s += white.Render(" ")
        case 1:
          s += white.Render(".")
        }
      }
    }
    s += "\n"
  }
  return s
}
