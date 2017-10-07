package headsUp

import "github.com/Hunrik/poker-player-goglig/leanpoker"

func getActivePlayers(players *[]leanpoker.Player) []leanpoker.Player {
	var activePlayers = []leanpoker.Player{}
	for _, player := range *players {
		if player.Status == "active" {
			activePlayers = append(activePlayers, player)
		}
	}
	return activePlayers
}
func calculate(gameState *leanpoker.Game) bool {
	activePlayers := getActivePlayers(&gameState.Players)
	return len(activePlayers) == 2
}
