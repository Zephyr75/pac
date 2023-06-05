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

type Teleport struct {
  A Point
  B Point
}


type direction int

const (
  Up    direction = 0
  Down  direction = 1
  Left  direction = 2
  Right direction = 3
  Idle  direction = 4
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
  Teleports []Teleport
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
      if g.PlayerPos.Y > 0 && g.GameMap[g.PlayerPos.Y - 1][g.PlayerPos.X] != 3 {
        g.PlayerDir = g.NextDir
      }
    case Down:
      if g.PlayerPos.Y < g.Height - 1 && g.GameMap[g.PlayerPos.Y + 1][g.PlayerPos.X] != 3 {
        g.PlayerDir = g.NextDir
      }
    case Left:
      if g.PlayerPos.X > 0 && g.GameMap[g.PlayerPos.Y][g.PlayerPos.X - 1] != 3 {
        g.PlayerDir = g.NextDir
      }
    case Right:
      if g.PlayerPos.X < g.Width - 1 && g.GameMap[g.PlayerPos.Y][g.PlayerPos.X + 1] != 3 {
        g.PlayerDir = g.NextDir
      }
    }
    switch g.PlayerDir {
    case Up:
      if g.PlayerPos.Y > 0 && g.GameMap[g.PlayerPos.Y - 1][g.PlayerPos.X] != 3 {
        g.PlayerPos.Y--
      }
    case Down:
      if g.PlayerPos.Y < g.Height - 1 && g.GameMap[g.PlayerPos.Y + 1][g.PlayerPos.X] != 3 {
        g.PlayerPos.Y++
      }
    case Left:
      if g.PlayerPos.X > 0 && g.GameMap[g.PlayerPos.Y][g.PlayerPos.X - 1] != 3 {
        g.PlayerPos.X--
      }
    case Right:
      if g.PlayerPos.X < g.Width - 1 && g.GameMap[g.PlayerPos.Y][g.PlayerPos.X + 1] != 3 {
        g.PlayerPos.X++
      }
    }
    if g.GameMap[g.PlayerPos.Y][g.PlayerPos.X] == 1 {
      g.Score++
      g.Dots--
      g.GameMap[g.PlayerPos.Y][g.PlayerPos.X] = 0
    }
    for i := 0; i < len(g.Teleports); i++ {
      if g.PlayerPos == g.Teleports[i].A {
        g.PlayerPos = g.Teleports[i].B
      } else if g.PlayerPos == g.Teleports[i].B {
        g.PlayerPos = g.Teleports[i].A
      }
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
  yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffff00"))
  white = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
  blue = lipgloss.NewStyle().Foreground(lipgloss.Color("#1919a6"))
  pink = lipgloss.NewStyle().Foreground(lipgloss.Color("#dea185"))
)

const (
  horiz = "━"
  vert = "┃"
  horizLeft = "╸"
  horizRight = "╺"
  vertUp = "╹"
  vertDown = "╻"
  topLeft = "┏"
  topRight = "┓"
  botLeft = "┗"
  botRight = "┛"
  tDown = "┳"
  tUp = "┻"
  tLeft = "┫"
  tRight = "┣"
  cross = "╋"

  dot = "•"
  bigDot = "●"
)

func (g Game) View() string {
  var s string
  if g.Dots <= 0 {
    s += "You win!\n"
    return s
  }
  s += "Score: " + white.Render(strconv.Itoa(g.Score)) + "\n"
  for y := 0; y < g.Height; y++ {
    for x := 0; x < g.Width; x++ {
      if x == g.PlayerPos.X && y == g.PlayerPos.Y {
        s += yellow.Render(g.PlayerChar)
      } else {
        switch g.GameMap[y][x] {
        case 0:
          s += white.Render(" ")
        case 1:
          s += pink.Render(dot)
        case 2:
          s += pink.Render(bigDot)
        case 3:
          left := 0
          if x > 0 {
            left = g.GameMap[y][x - 1]
          }
          right := 0
          if x < g.Width - 1 {
            right = g.GameMap[y][x + 1]
          }
          up := 0
          if y > 0 {
            up = g.GameMap[y - 1][x]
          }
          down := 0
          if y < g.Height - 1 {
            down = g.GameMap[y + 1][x]
          }

          if left == 3 && right == 3 && up == 3 && down == 3 {
            s += blue.Render(cross)
          } else if left == 3 && right == 3 && up == 3 {
            s += blue.Render(tUp)
          } else if left == 3 && right == 3 && down == 3 {
            s += blue.Render(tDown)
          } else if left == 3 && up == 3 && down == 3 {
            s += blue.Render(tLeft)
          } else if right == 3 && up == 3 && down == 3 {
            s += blue.Render(tRight)
          } else if left == 3 && up == 3 {
            s += blue.Render(botRight)
          } else if left == 3 && down == 3 {
            s += blue.Render(topRight)
          } else if right == 3 && up == 3 {
            s += blue.Render(botLeft)
          } else if right == 3 && down == 3 {
            s += blue.Render(topLeft)
          } else if up == 3 && down == 3 {      
            s += blue.Render(vert)
          } else if left == 3 && right == 3 {
            s += blue.Render(horiz)
          } else if left == 3 {
            s += blue.Render(horizLeft)
          } else if right == 3 {
            s += blue.Render(horizRight)
          } else if up == 3 {
            s += blue.Render(vertUp)
          } else if down == 3 {
            s += blue.Render(vertDown)
          }
        }
      }
    }
    s += "\n"
  }
  s += yellow.Render("CCC")
  return s
}
