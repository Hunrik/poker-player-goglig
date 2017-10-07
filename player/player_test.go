package player

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/Hunrik/poker-player-goglig/leanpoker"
	"github.com/tj/assert"
)

func getFixture(path string) (game *leanpoker.Game, err error) {
	raw, err := ioutil.ReadFile(path)
	game = &leanpoker.Game{}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err = json.Unmarshal(raw, game); err != nil {
		log.Printf("Error parsing game state: %s", err)
		return nil, err
	}

	return game, nil
}
func TestMyPlayer(t *testing.T) {
	gameState, _ := getFixture("./fixture-1.json")
	ret := GetMyPlayer(gameState)
	assert.Equal(t, 1, ret.Id)
}
