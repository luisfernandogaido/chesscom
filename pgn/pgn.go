package pgn

import (
	"time"
	"strings"
	"strconv"
	"path/filepath"
	"os"
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
	//	Mon Jan 2 15:04:05 -0700 MST 2006
	ret := "[Event \"" + g.Event + "\"]\n" +
		"[Site \"" + g.Site + "\"]\n" +
		"[Date \"" + g.Date.Format("2006.01.02") + "\"]\n" +
		"[Round \"" + g.Round + "\"]\n" +
		"[White \"" + g.White + "\"]\n" +
		"[Black \"" + g.Black + "\"]\n" +
		"[Result \"" + g.Result + "\"]\n" +
		"[WhiteElo \"" + strconv.Itoa(g.WhiteElo) + "\"]\n" +
		"[BlackElo \"" + strconv.Itoa(g.BlackElo) + "\"]\n" +
		"[TimeControl \"" + g.TimeControl + "\"]\n" +
		"[Termination \"" + g.Termination + "\"]\n" +
		"[StartTime \"" + g.StartTime.Format("15:04:05") + "\"]\n" +
		"[EndDate \"" + g.EndDate.Format("2006.01.02") + "\"]\n" +
		"[EndTime \"" + g.EndTime.Format("15:04:05") + "\"]\n" +
		"[Link \"" + g.Link + "\"]\n" +
		"\n" +
		g.Moves + "\n" +
		"\n"
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
	fileName := filepath.Join(folder, time.Now().Format("20060102-150405")+".pgn")
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()
	for _, g := range games {
		_, err = f.WriteString(g.String())
		if err != nil {
			return "", err
		}
	}
	return fileName, nil
}
