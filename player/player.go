package player

import "github.com/Hunrik/poker-player-goglig/leanpoker"

const VERSION = "High rollers 20"

func BetRequest(state *leanpoker.Game) int {
	return 20
}

func Showdown(state *leanpoker.Game) {

}

func Version() string {
	return VERSION
}
