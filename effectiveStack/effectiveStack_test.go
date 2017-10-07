package effectiveStack

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
func TestEffStack(t *testing.T) {
	t.Run("Eff stack 1", func(t *testing.T) {
		gameState, err := getFixture("./fixture-1.json")
		assert.NoError(t, err, "read body")
		fmt.Println(gameState)
		resp := EffStack(gameState)
		assert.Equal(t, float64(1010+320)/float64(20), resp)
	})
	t.Run("Eff stack 2", func(t *testing.T) {
		gameState, err := getFixture("./fixture-4.json")
		assert.NoError(t, err, "read body")

		resp := EffStack(gameState)
		assert.Equal(t, 43.5, resp)
	})
}
