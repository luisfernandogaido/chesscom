package main

import (
	"github.com/luisfernandogaido/chesscom/api"
	"log"
	"os"
	"fmt"
)

func main() {
	archives, err := api.Archives("luisfernandogaido")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(archives))
	f, err := os.OpenFile("./luisfernandogaido.pgn", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for k, a := range archives {
		bytes, err := api.MultiGamePgn(a)
		if err != nil {
			log.Fatal(err)
		}
		if _, err = f.Write(bytes); err != nil {
			log.Fatal(err)
		}
		fmt.Println(k)
	}
}
