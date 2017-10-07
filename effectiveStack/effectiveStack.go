package effectiveStack

import "github.com/Hunrik/poker-player-goglig/leanpoker"

func EffStack(gameState *leanpoker.Game) float64 {
	bigBlind := (2 * gameState.SmallBlind)
	myPlayerBet := gameState.Players[gameState.InAction].Bet
	myPlayerStack := gameState.Players[gameState.InAction].Stack
	var myPlayerSum float64 = float64(myPlayerBet+myPlayerStack) / float64(bigBlind)

	activePlayers := 0
	var otherPlayersSumStack float64 = 0

	for _, player := range gameState.Players {
		if player.Id != gameState.InAction && player.Status == "active" {
			activePlayers++
			stack := player.Stack
			bet := player.Bet
			otherPlayersSumStack += float64(stack+bet) / float64(bigBlind)
		}
	}

	otherPlayerAverageStack := float64(otherPlayersSumStack) / float64(activePlayers)

	if myPlayerSum < otherPlayerAverageStack {
		return myPlayerSum
	}
	return otherPlayerAverageStack
}
