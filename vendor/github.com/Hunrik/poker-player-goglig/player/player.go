package player

import "github.com/lean-poker/poker-player-go/leanpoker"

const VERSION = "0.0.1 Player"

func BetRequest(state *leanpoker.Game) int {
	return 0
}

func Showdown(state *leanpoker.Game) {

}

func Version() string {
	return VERSION
}
