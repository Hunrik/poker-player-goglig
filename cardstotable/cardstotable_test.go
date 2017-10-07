package cardstotable

import (
	"testing"

	"github.com/Hunrik/poker-player-goglig/leanpoker"
	"github.com/tj/assert"
)

func TestCardsToTableJ(t *testing.T) {
	ret := ToNum("J")
	assert.Equal(t, 11, ret)
}

func TestCardsToTableK(t *testing.T) {
	ret := ToNum("K")
	assert.Equal(t, 13, ret)
}

func TestCardsToTableA(t *testing.T) {
	ret := ToNum("A")
	assert.Equal(t, 14, ret)
}

func TestCardsToTable8(t *testing.T) {
	ret := ToNum("8")
	assert.Equal(t, 8, ret)
}

func TestFromatRankT(t *testing.T) {
	ret := FormatRank("10")
	assert.Equal(t, "T", ret)
}

func TestFromatRankDef(t *testing.T) {
	ret := FormatRank("J")
	assert.Equal(t, "J", ret)
}

func TestConvert2(t *testing.T) {
	cards := [2]leanpoker.Card{{"4", "spades"}, {"5", "spades"}}
	assert.Equal(t, Convert(cards), "54s")

	cards = [2]leanpoker.Card{{"4", "spades"}, {"4", "spades"}}
	assert.Equal(t, Convert(cards), "44")

	cards = [2]leanpoker.Card{{"4", "spades"}, {"5", "hearts"}}
	assert.Equal(t, Convert(cards), "54o")

	cards = [2]leanpoker.Card{{"10", "spades"}, {"4", "hearts"}}
	assert.Equal(t, Convert(cards), "T4o")
}

func TestPercentage(t *testing.T) {
	cards := [2]leanpoker.Card{{"A", "spades"}, {"A", "spades"}}
	assert.Equal(t, GetPercentage(cards), 0.5)

	cards = [2]leanpoker.Card{{"J", "spades"}, {"K", "spades"}}
	assert.Equal(t, GetPercentage(cards), 12.2)

	cards = [2]leanpoker.Card{{"A", "spades"}, {"Q", "clubs"}}
	assert.Equal(t, GetPercentage(cards), 4.7)

	cards = [2]leanpoker.Card{{"J", "spades"}, {"J", "clubs"}}
	assert.Equal(t, GetPercentage(cards), float64(3))

	cards = [2]leanpoker.Card{{"8", "spades"}, {"5", "clubs"}}
	assert.Equal(t, GetPercentage(cards), 78.3)

}

/*
test('cards 2 table percentage should work', (t) => {
  let cards = [{rank: 'A', suit: 'spades'}, {rank: 'A', suit: 'spades'}]
  t.equal(cards2Table.getPercentage(cards), 0.5)

  cards = [{rank: 'J', suit: 'spades'}, {rank: 'K', suit: 'spades'}]
  t.equal(cards2Table.getPercentage(cards), 12.2)

  cards = [{rank: 'A', suit: 'spades'}, {rank: 'Q', suit: 'clubs'}]
  t.equal(cards2Table.getPercentage(cards), 4.7)

  cards = [{rank: 'J', suit: 'spades'}, {rank: 'J', suit: 'clubs'}]
  t.equal(cards2Table.getPercentage(cards), 3)

  cards = [{rank: '8', suit: 'spades'}, {rank: '5', suit: 'clubs'}]
  t.equal(cards2Table.getPercentage(cards), 78.3)

  t.end()
})

})
*/
