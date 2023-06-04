package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
  "pac/load"
)


func main() {
  p := tea.NewProgram(load.LoadGame(1))
  if err := p.Start(); err != nil {
    fmt.Println("Error running program:", err)
    os.Exit(1)
  }
}
