package player

import (
	"math/rand"

	"github.com/Hunrik/poker-player-goglig/leanpoker"
)

const VERSION = "High rollers lol"

func BetRequest(state *leanpoker.Game) int {
	return rand.Intn(100) + 100
}

func Showdown(state *leanpoker.Game) {

}

func Version() string {
	return VERSION
}
