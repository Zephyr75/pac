package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
  "pac/game" 
  "pac/load"
)




func main() {
  width := 30
  height := 10
  p := tea.NewProgram(game.InitGame(width, height))
  if err := p.Start(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }
}
