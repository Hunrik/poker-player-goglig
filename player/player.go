package player

import (
	"fmt"

	"math/rand"

	"github.com/Hunrik/poker-player-goglig/effectiveStack"
	"github.com/Hunrik/poker-player-goglig/leanpoker"
)

const VERSION = "High rollers lol"

func BetRequest(state *leanpoker.Game) int {
	effectiveStack := effectiveStack.EffStack(state)
	fmt.Println("Effective stack", effectiveStack)
	return rand.Intn(100) + 100
}

func Showdown(state *leanpoker.Game) {

}

func Version() string {
	return VERSION
}
