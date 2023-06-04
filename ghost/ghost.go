package ghost

type ghostType int

const (
  blinky ghostType = 0
  inky   ghostType = 1
  pinky  ghostType = 2
  clyde  ghostType = 3
)

type Ghost struct {
  pos Point
  char string
  ghostType ghostType
}


func NewGame() Game {
  return Game{}
}
