package load

import (
	"os"
  "bufio"
	"log"
  "strconv"

)

func LoadGame(level int) main.Game {
  name := "level_" + strconv.Itoa(level) + ".txt"
  file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

  gameMap := make([][]int, 0)
  playerPos := game.Point{X: 0, Y: 0}

  lineNbr := 0
	for scanner.Scan() {
    for i, c := range scanner.Text() {
      line := make([]int, 0)
      switch c {
      case '.':
        line = append(line, 1)
      case '#':
        line = append(line, 1)
      case 'C':
        playerPos = game.Point{X: i, Y: lineNbr}
      default:
        line = append(line, 0)
      }
      gameMap = append(gameMap, line)
    }
    lineNbr++
  }

  game := game.Game {
    Width: len(gameMap[0]),
    Height: len(gameMap),
    PlayerPos: playerPos,
    PlayerChar: "C",
    GameMap: gameMap,
    Ghosts: make([]game.Ghost, 0),
  }

  return game
}
