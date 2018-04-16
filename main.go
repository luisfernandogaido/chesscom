package main

import (
	"io/ioutil"
	"log"
	"github.com/luisfernandogaido/chesscom/api"
	"github.com/luisfernandogaido/chesscom/pgn"
	"fmt"
)

const (
	user     = "luisfernandogaido"
	folder   = "C:\\Users\\lfgai\\Desktop"
	lastFile = "C:\\Users\\lfgai\\Desktop\\last.txt"
)

func main() {
	last, err := Last()
	if err != nil {
		log.Fatal(err)
	}
	games, err := AllAfter(last)
	games = pgn.Reverse(games)
	fileName, err := pgn.Save(games, folder)
	if err != nil {
		log.Fatal(err)
	}
	if len(games) == 0 {
		fmt.Println("Nenhuma partida nova.")
		return
	}
	fmt.Printf("%v salvo com %v partidas.\n", fileName, len(games))
	last = games[len(games)-1].Link
	err = SaveLast(last)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v salvo em last.txt\n", last)
}

func Last() (string, error) {
	bytes, err := ioutil.ReadFile(lastFile)
	return string(bytes), err
}

func SaveLast(l string) (error) {
	return ioutil.WriteFile(lastFile, []byte(l), 0666)
}

func AllAfter(last string) ([]pgn.Game, error) {
	archives, err := api.Archives(user)
	if err != nil {
		return nil, err
	}
	all := make([]pgn.Game, 0)
Loop:
	for i := len(archives) - 1; i >= 0; i-- {
		arc := archives[i]
		bytes, err := api.MultiGamePgn(arc)
		if err != nil {
			return nil, err
		}
		gamesArc, err := pgn.Parse(bytes)
		if err != nil {
			return nil, err
		}
		for _, g := range gamesArc {
			if g.Link == last {
				break Loop
			}
			all = append(all, g)
		}
	}
	return all, nil
}
