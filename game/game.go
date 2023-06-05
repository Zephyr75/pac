package game

import (
  "pac/ghost"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
  "time"
  "strconv"
)

type Point struct {
  X int
  Y int
}


type direction int

const (
  Up    direction = 0
  Down  direction = 1
  Left  direction = 2
  Right direction = 3
)

type Game struct {
  Width int
  Height int
  PlayerPos Point
  PlayerDir direction
  NextDir direction
  PlayerChar string
  GameMap [][]int
  Score int
  Dots int
  Ghosts []ghost.Ghost
}

type updateMsg int

func updateGame() tea.Msg {
  time.Sleep(100 * time.Millisecond)
  return updateMsg(0)
}


func (g Game) Init() tea.Cmd {
  // g.ghosts = append(g.ghosts, ghost{point{g.Width - 2, g.Height - 2}, "B", blinky})
  // g.ghosts = append(g.ghosts, ghost{point{g.Width - 2, 1}, "I", inky})
  // g.ghosts = append(g.ghosts, ghost{point{1, g.Height - 2}, "P", pinky})
  // g.ghosts = append(g.ghosts, ghost{point{1, 1}, "Y", clyde})
  return updateGame
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "up", "k":
      g.NextDir = Up
    case "down", "j":
      g.NextDir = Down
    case "left", "h":
      g.NextDir = Left
    case "right", "l":
      g.NextDir = Right
    case "ctrl+c":
      return g, tea.Quit
    }
  case updateMsg:
    switch g.NextDir {
    case Up:
      if g.PlayerPos.Y > 0 && g.GameMap[g.PlayerPos.Y - 1][g.PlayerPos.X] != 2 {
        g.PlayerDir = g.NextDir
      }
    case Down:
      if g.PlayerPos.Y < g.Height - 1 && g.GameMap[g.PlayerPos.Y + 1][g.PlayerPos.X] != 2 {
        g.PlayerDir = g.NextDir
      }
    case Left:
      if g.PlayerPos.X > 0 && g.GameMap[g.PlayerPos.Y][g.PlayerPos.X - 1] != 2 {
        g.PlayerDir = g.NextDir
      }
    case Right:
      if g.PlayerPos.X < g.Width - 1 && g.GameMap[g.PlayerPos.Y][g.PlayerPos.X + 1] != 2 {
        g.PlayerDir = g.NextDir
      }
    }
    switch g.PlayerDir {
    case Up:
      if g.PlayerPos.Y > 0 && g.GameMap[g.PlayerPos.Y - 1][g.PlayerPos.X] != 2 {
        g.PlayerPos.Y--
      }
    case Down:
      if g.PlayerPos.Y < g.Height - 1 && g.GameMap[g.PlayerPos.Y + 1][g.PlayerPos.X] != 2 {
        g.PlayerPos.Y++
      }
    case Left:
      if g.PlayerPos.X > 0 && g.GameMap[g.PlayerPos.Y][g.PlayerPos.X - 1] != 2 {
        g.PlayerPos.X--
      }
    case Right:
      if g.PlayerPos.X < g.Width - 1 && g.GameMap[g.PlayerPos.Y][g.PlayerPos.X + 1] != 2 {
        g.PlayerPos.X++
      }
    }
    if g.GameMap[g.PlayerPos.Y][g.PlayerPos.X] == 1 {
      g.Score++
      g.Dots--
      g.GameMap[g.PlayerPos.Y][g.PlayerPos.X] = 0
    }
    if g.PlayerChar == "C" { 
      g.PlayerChar = "c"
    } else {
      g.PlayerChar = "C"
    }
    return g, updateGame
  }
  return g, nil
}




var (
  yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffeb3b"))
  white = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
)

func (g Game) View() string {
  var s string
  if g.Dots <= 0 {
    s += "You win!\n"
    return s
  }
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
  s += "Score: " + white.Render(strconv.Itoa(g.Score)) + "\n"
  s += "Dots: " + white.Render(strconv.Itoa(g.Dots)) + "\n"
  s += "Current direction: " + white.Render(strconv.Itoa(int(g.PlayerDir))) + "\n"
  s += "Next direction: " + white.Render(strconv.Itoa(int(g.NextDir))) + "\n"
  return s
}
