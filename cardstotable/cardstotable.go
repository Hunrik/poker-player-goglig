package cardstotable

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/Hunrik/poker-player-goglig/leanpoker"
)

func ToNum(char string) int {
	switch char {
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		x, _ := strconv.Atoi(char)
		return x

	}
}

func FormatRank(rank string) string {
	switch rank {
	case "10":
		return "T"
	default:
		return rank
	}
}

func Convert(cards [2]leanpoker.Card) string {
	var tableStr = ""
	firstCardNum := ToNum(cards[0].Rank)
	secondCardNum := ToNum(cards[1].Rank)

	firstCardString := FormatRank(cards[0].Rank)
	secondCardString := FormatRank(cards[1].Rank)

	if firstCardNum > secondCardNum {
		tableStr += firstCardString + secondCardString
	} else {
		tableStr += secondCardString + firstCardString
	}

	if firstCardNum == secondCardNum {
		tableStr += "" // do nothing
	} else if cards[0].Suit == cards[1].Suit {
		tableStr += "s"
	} else {
		tableStr += "o"
	}

	return tableStr

}

// type PreflopTable struct {
// 	// Rank of the card. Possible values are numbers 2-10 and J,Q,K,A
// 	Cards string `json:"cards"`

// 	// Suit of the card. Possible values are: clubs,spades,hearts,diamonds
// 	Percentage float64 `json:"percentage"`
// }

func getFixture(path string) (table *[]PreflopTable, err error) {
	raw, err := ioutil.ReadFile(path)
	table = &[]PreflopTable{}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err = json.Unmarshal(raw, table); err != nil {
		log.Printf("Error parsing game state: %s", err)
		return nil, err
	}

	return table, nil
}

func GetPercentage(cards [2]leanpoker.Card) float64 {
	cardStr := Convert(cards)
	// preFlopTable, _ := getFixture("./preflop-table.json")
	// fmt.Println("%v", preFlopTable)
	// for i := 0; i < len(*preFlopTable); i++ {
	// 	if preFlopTable[i].Cards == cardStr {
	// 		return V.Percentage
	// 	}
	// 	// S[i].mark = "S"
	// }
	for _, v := range PreflopTableArray {
		if v.Cards == cardStr {
			return v.Percentage
		}
	}
	return 0
	// const record = _.findWhere(preflopTable, {cards: cardStr})
	// return Number(record.percentage)
}
