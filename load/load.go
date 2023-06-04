package load

import (
	"os"
  "bufio"
	"log"
  "strconv"

  "pac/game" 
)

func LoadGame(level int) game.Game {
  name := "load/levels/level_" + strconv.Itoa(level) + ".txt"
  file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

  gameMap := make([][]int, 0)
  playerPos := game.Point{X: 0, Y: 0}

  width := 0
  height := 0

  lineNbr := 0
	for scanner.Scan() {
    if (len(scanner.Text()) == 0) {
      break
    }
    width = len(scanner.Text())
    line := make([]int, 0)
    for i, c := range scanner.Text() {
      switch c {
      case '.':
        line = append(line, 1)
      case '#':
        line = append(line, 2)
      case 'C':
        playerPos = game.Point{X: i, Y: lineNbr}
        line = append(line, 1)
      default:
        line = append(line, 0)
      }
    }
    gameMap = append(gameMap, line)
    lineNbr++
  }
  height = lineNbr

  game := game.Game {
    Width: width,
    Height: height,
    PlayerPos: playerPos,
    PlayerChar: "C",
    GameMap: gameMap,
    // Ghosts: make([]game.Ghost, 0),
  }

  return game
}
