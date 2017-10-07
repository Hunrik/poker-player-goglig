package headsup

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

func TestHeadsUp(t *testing.T) {
	t.Run("heads up should be false", func(t *testing.T) {
		gameState, err := getFixture("./fixture-4.json")
		assert.NoError(t, err, "read body")

		resp := calculate(gameState)
		assert.Equal(t, false, resp)
	})
	t.Run("Eff stack 2", func(t *testing.T) {
		gameState, err := getFixture("./fixture-1.json")
		assert.NoError(t, err, "read body")

		resp := calculate(gameState)
		assert.Equal(t, true, resp)
	})
}

/*test('heads up should be false', (t) => {
  const isHeadsUp = headsUp.calculate(gameStateFixture4)
  t.equal(isHeadsUp, false)
  t.end()
})

test('heads up should be true', (t) => {
  const isHeadsUp = headsUp.calculate(gameStateFixture)
  t.equal(isHeadsUp, true)
  t.end()
})
*/
