package player

import (
	"github.com/Hunrik/poker-player-goglig/cardstotable"
	"github.com/Hunrik/poker-player-goglig/effectiveStack"
	"github.com/Hunrik/poker-player-goglig/headsup"
	"github.com/Hunrik/poker-player-goglig/leanpoker"
)

const VERSION = "High rollers coming"

func GetMyPlayer(state *leanpoker.Game) leanpoker.Player {
	for _, v := range state.Players {
		if v.Id == state.InAction {
			return v
		}
	}
	return leanpoker.Player{}
}

func isRaised(state *leanpoker.Game) bool {
	return state.Pot > (state.SmallBlind * 5)
}

func BetRequest(state *leanpoker.Game) int {
	myPlayer := GetMyPlayer(state)
	cards := myPlayer.HoleCards
	percentage := cardstotable.GetPercentage(cards)
	raised := isRaised(state)

	isHeadsup := headsup.Calculate(state)
	effectiveStack := effectiveStack.EffStack(state)

	betValue := 0
	// If not heads up
	if effectiveStack <= 3 {
		if !raised && percentage < 81 {
			betValue = myPlayer.Stack
		} else if percentage < 48 {
			betValue = myPlayer.Stack
		}
	} else if effectiveStack > 3 && effectiveStack <= 10 {
		if !raised && percentage < 31 {
			betValue = myPlayer.Stack
		} else if percentage < 15 {
			betValue = myPlayer.Stack
		}
	} else if effectiveStack > 10 && effectiveStack <= 20 {
		if !raised && percentage < 21 {
			betValue = myPlayer.Stack
		} else if percentage < 13 {
			betValue = myPlayer.Stack
		}
	} else if effectiveStack > 20 && effectiveStack <= 50 {
		if !raised && percentage < 8 {
			betValue = myPlayer.Stack
		} else if percentage < 6 {
			betValue = myPlayer.Stack
		}
	} else {
		if !raised && percentage < 4 {
			betValue = myPlayer.Stack
		} else if percentage < 3.1 {
			betValue = myPlayer.Stack
		}
	}

	// If headsup
	// headsUpCall := false
	if isHeadsup {
		if state.Pot <= state.SmallBlind*4 {
			if effectiveStack <= 3 {
				betValue = myPlayer.Stack
			} else if effectiveStack > 3 && effectiveStack <= 6 {
				if percentage < 78.7 {
					betValue = myPlayer.Stack
				}
			} else if effectiveStack > 6 && effectiveStack <= 10 {
				if percentage < 61 {
					betValue = myPlayer.Stack
				}
			} else if effectiveStack > 10 && effectiveStack <= 15 {
				if percentage < 50 {
					betValue = myPlayer.Stack
				}
			} else if effectiveStack > 15 && effectiveStack <= 25 {
				if percentage < 26 {
					betValue = myPlayer.Stack
				}
			} else if effectiveStack > 25 && effectiveStack <= 50 {
				if percentage < 18 {
					betValue = myPlayer.Stack
				} else if percentage < 3.1 {
					betValue = myPlayer.Stack
				}
			}
		} else {
			// headsUpCall = true
			if effectiveStack <= 3 {
				if percentage < 77 {
					betValue = myPlayer.Stack
				}
			} else if effectiveStack > 3 && effectiveStack <= 6 {
				if percentage < 49 {
					betValue = myPlayer.Stack
				}
			} else if effectiveStack > 6 && effectiveStack <= 10 {
				if percentage < 42 {
					betValue = myPlayer.Stack
				}
			} else if effectiveStack > 10 && effectiveStack <= 15 {
				if percentage < 29 {
					betValue = myPlayer.Stack
				}
			} else if effectiveStack > 15 && effectiveStack <= 25 {
				if percentage < 22 {
					betValue = myPlayer.Stack
				}
			} else if effectiveStack > 25 && effectiveStack <= 50 {
				if percentage < 13 {
					betValue = myPlayer.Stack
				} else if percentage < 3.1 {
					betValue = myPlayer.Stack
				}
			}
		}
	}

	return betValue

}

func Showdown(state *leanpoker.Game) {

}

func Version() string {
	return VERSION
}
