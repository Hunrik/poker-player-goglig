package player

import (
	"fmt"

	"github.com/Hunrik/poker-player-goglig/effectiveStack"
	"github.com/Hunrik/poker-player-goglig/leanpoker"
)

const VERSION = "High rollers 20"

func BetRequest(state *leanpoker.Game) int {
	effectiveStack = effectiveStack.EffStack(state)
	fmt.Println("Effective stack", effectiveStack)
	return 20
}

func Showdown(state *leanpoker.Game) {

}

func Version() string {
	return VERSION
}
