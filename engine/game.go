package engine

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jonasdacruz/lighthouses_aicontest/coms"
)

const (
	timeoutPlayerToAnswer = time.Millisecond * 100
	gameTurns             = 10
)

type Game struct {
	gameStartAt time.Time
	mapMatrix   []*coms.MapRow
	lighthouses []*coms.Lighthouse
	players     []Player
}

func loadMap() ([]*coms.MapRow, []*coms.Lighthouse) {

	mapFile := filepath.Join(os.Getenv("PWD"), "maps", "island.txt")
	file, err := os.Open(mapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	var mapM []*coms.MapRow
	var lighthouses []*coms.Lighthouse

	for x := range fileLines {

		var row coms.MapRow
		for y, val := range fileLines[len(fileLines)-x-1] {

			switch c := string(val); c {
			case "#":
				row.Row = append(row.Row, 0)

			case "!":
				row.Row = append(row.Row, 1)
				pos := coms.Position{
					X: int32(x),
					Y: int32(y),
				}
				lh := coms.Lighthouse{
					Position: &pos,
				}
				lighthouses = append(lighthouses, &lh)
			default:
				row.Row = append(row.Row, 1)
			}

		}
		mapM = append(mapM, &row)
	}
	// DEBUG map
	// for i, ri := range mapM {
	// 	fmt.Printf("row %d -> %+v\n", i, ri)
	// }
	// DEBUG lighthouses
	// fmt.Println(len(lighthouses))
	// for _, l := range lighthouses {
	// 	fmt.Println(l.Position.GetX(), l.Position.GetY())
	// }

	w := len(mapM[0].Row)
	h := len(mapM)

	if w != h {
		panic("Island must be a square")
	}
	for i, mapR := range mapM {
		if i == 0 || i == len(mapM)-1 {
			for _, c := range mapR.Row {
				if c != 0 {
					panic("Map border must not be part of island")
				}
			}
		} else {
			if mapR.Row[0] != 0 || mapR.Row[len(mapR.Row)-1] != 0 {
				panic("Map border must not be part of island")

			}
		}
	}
	return mapM, lighthouses
}

func NewGame() *Game {
	mapM, lighthouses := loadMap()
	return &Game{
		gameStartAt: time.Now(),
		mapMatrix:   mapM,
		lighthouses: lighthouses,
		players:     []Player{},
	}
}

func (e *Game) AddNewPlayer(np Player) error {
	e.players = append(e.players, np)
	return nil
}

func (e *Game) CreateInitialState(p Player) *coms.NewPlayerInitialState {
	// TODO: implement the logic to create the initial state of the game
	lighthouseExample := &coms.Lighthouse{Position: &coms.Position{X: 1, Y: 1}}
	npst := &coms.NewPlayerInitialState{
		PlayerNum:   0,
		PlayerCount: 2,
		Position:    &coms.Position{X: 1, Y: 2},
		Map:         []*coms.MapRow{{Row: []int32{0, 0, 0, 0, 0}}},
		Lighthouses: []*coms.Lighthouse{lighthouseExample},
	}
	return npst
}

func (e *Game) StartGame() {
	e.gameStartAt = time.Now()

	viewExample := &coms.MapRow{
		Row: []int32{0, 0, 0, 0, 0},
	}
	lighthouseExample := &coms.Lighthouse{
		Position: &coms.Position{
			X: 1,
			Y: 1,
		},
		Energy: 100,
	}

	newTurnExample := &coms.NewTurn{
		Position: &coms.Position{
			X: 1,
			Y: 1,
		},
		Score:       10,
		Energy:      100,
		View:        []*coms.MapRow{viewExample},
		Lighthouses: []*coms.Lighthouse{lighthouseExample},
	}

	for i := 0; i < gameTurns; i++ {
		// TODO: calc new turn state
		for _, p := range e.players {
			gc := &GameClient{}
			comPlayer := &coms.NewPlayer{
				Name:          p.Name,
				ServerAddress: p.ServerAddress,
			}
			// TODO: calc the new turn with the last state of the game
			na := gc.requestTurn(comPlayer, newTurnExample)
			_ = na
			// TODO: apply the action to change the state of the game
		}
		// TODO: calc players scores
	}
}
