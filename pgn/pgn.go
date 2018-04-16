package pgn

import (
	"time"
	"strings"
	"strconv"
	"path/filepath"
)

type Game struct {
	Event       string
	Site        string
	Date        time.Time
	Round       string
	White       string
	Black       string
	Result      string
	WhiteElo    int
	BlackElo    int
	TimeControl string
	Termination string
	StartTime   time.Time
	EndDate     time.Time
	EndTime     time.Time
	Link        string
	Moves       string
}

func (g *Game) String() string {
	ret := `[Event "Live Chess"]`
	return ret
}

func Parse(b []byte) ([]Game, error) {
	ll := strings.Split(string(b), "\n")
	games := make([]Game, 0)
	var game Game
	for _, l := range ll {
		if l == "" {
			continue
		}
		if strings.HasPrefix(l, "[Event ") {
			game = Game{}
		}
		var err error
		if strings.HasPrefix(l, "[") {
			atributo, valor := attr(l)
			switch atributo {
			case "Event":
				game.Event = valor
			case "Site":
				game.Site = valor
			case "Date":
				game.Date, err = time.Parse("2006.01.02", valor)
				if err != nil {
					return nil, err
				}
			case "Round":
				game.Round = valor
			case "White":
				game.White = valor
			case "Black":
				game.Black = valor
			case "Result":
				game.Result = valor
			case "WhiteElo":
				game.WhiteElo, err = strconv.Atoi(valor)
				if err != nil {
					return nil, err
				}
			case "BlackElo":
				game.BlackElo, err = strconv.Atoi(valor)
				if err != nil {
					return nil, err
				}
			case "TimeControl":
				game.TimeControl = valor
			case "Termination":
				game.Termination = valor
			case "StartTime":
				game.StartTime, err = time.Parse("15:04:05", valor)
				if err != nil {
					return nil, err
				}
			case "EndDate":
				game.EndDate, err = time.Parse("2006.01.02", valor)
				if err != nil {
					return nil, err
				}
			case "EndTime":
				game.EndTime, err = time.Parse("15:04:05", valor)
				if err != nil {
					return nil, err
				}
			case "Link":
				game.Link = valor
			}
			continue
		}
		game.Moves = l
		games = append(games, game)
	}
	return games, nil
}

func attr(l string) (string, string) {
	atributo := l[1:strings.Index(l, " ")]
	valor := l[strings.Index(l, "\"")+1 : strings.LastIndex(l, "\"")]
	return atributo, valor
}

func Reverse(games []Game) []Game {
	rev := make([]Game, 0)
	for i := len(games) - 1; i >= 0; i-- {
		rev = append(rev, games[i])
	}
	return rev
}

func Save(games []Game, folder string) (string, error) {
	//	Mon Jan 2 15:04:05 -0700 MST 2006
	fileName := filepath.Join(folder, time.Now().Format("20060102-150405")+".pgn")

	return "", nil
}
