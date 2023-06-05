package load

import (
	"os"
  "bufio"
	"log"
  "strconv"

  "pac/game" 
  "pac/utils"
  "pac/ghost"
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
  playerPos := utils.Point{X: 0, Y: 0}

  width := 0
  height := 0
  dots := 0
  mapTeleports := make(map[int][]utils.Point)
  ghosts := make([]ghost.Ghost, 0)

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
        dots++
      case 'o':
        line = append(line, 2)
      case '#':
        line = append(line, 3)
      case 'c':
        playerPos = utils.Point{X: i, Y: lineNbr}
        line = append(line, 1)
      case 'B':
        ghosts = append(ghosts, ghost.Ghost{
          Pos: utils.Point{X: i, Y: lineNbr}, 
          GhostType: ghost.Blinky,
        })
        line = append(line, 0)
      case 'I':
        ghosts = append(ghosts, ghost.Ghost{
          Pos: utils.Point{X: i, Y: lineNbr},
          GhostType: ghost.Inky,
        })
        line = append(line, 0)
      case 'P':
        ghosts = append(ghosts, ghost.Ghost{
          Pos: utils.Point{X: i, Y: lineNbr},
          GhostType: ghost.Pinky,
        })
        line = append(line, 0)
      case 'C':
        ghosts = append(ghosts, ghost.Ghost{
          Pos: utils.Point{X: i, Y: lineNbr},
          GhostType: ghost.Clyde,
        })
        line = append(line, 0)
      case '1':
        mapTeleports[1] = append(mapTeleports[1], utils.Point{X: i, Y: lineNbr})
        line = append(line, 0)
      case '2':
        mapTeleports[2] = append(mapTeleports[2], utils.Point{X: i, Y: lineNbr})
        line = append(line, 0)
      case '3':
        mapTeleports[3] = append(mapTeleports[3], utils.Point{X: i, Y: lineNbr})
        line = append(line, 0)
      default:
        line = append(line, 0)
      }
    }
    gameMap = append(gameMap, line)
    lineNbr++
  }
  height = lineNbr

  teleports := make([]game.Teleport, 0)
  for i := 0; i < len(mapTeleports[1]); i += 2 {
    teleports = append(teleports, game.Teleport{
      A: mapTeleports[1][i],
      B: mapTeleports[1][i + 1],
    })
  }

  game := game.Game {
    Width: width,
    Height: height,
    PlayerPos: playerPos,
    PlayerChar: "C",
    PlayerDir: game.Idle,
    GameMap: gameMap,
    Dots: dots,
    Teleports: teleports,
    Ghosts: ghosts,
  }

  return game
}
