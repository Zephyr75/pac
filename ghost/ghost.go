package ghost

import (
  "pac/utils"
)

type ghostType int

const (
  Blinky ghostType = 0
  Inky   ghostType = 1
  Pinky  ghostType = 2
  Clyde  ghostType = 3
)

type Ghost struct {
  Pos utils.Point
  GhostType ghostType
}

