package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseUrl = "https://api.chess.com/pub"
)

type archives struct {
	Archives []string `json:"archives"`
}

func Archives(user string) ([]string, error) {
	res, err := http.Get(fmt.Sprintf("%v/player/%v/games/archives", baseUrl, user))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var a archives
	err = json.Unmarshal(bytes, &a)
	return a.Archives, err
}

func MultiGamePgn(archive string) ([]byte, error) {
	res, err := http.Get(fmt.Sprintf("%v/pgn", archive))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
