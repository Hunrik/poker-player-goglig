package cardstotable

type PreflopTable struct {
	Cards      string  `json:"cards"`
	Percentage float64 `json:"percentage"`
}

var PreflopTableArray = []PreflopTable{
	PreflopTable{
		Cards:      "AA",
		Percentage: 0.5,
	},
	PreflopTable{
		Cards:      "KK",
		Percentage: 0.9,
	},
	PreflopTable{
		Cards:      "AKs",
		Percentage: 1.2,
	},
	PreflopTable{
		Cards:      "QQ",
		Percentage: 1.7,
	},
	PreflopTable{
		Cards:      "AKo",
		Percentage: 2.6,
	},
	PreflopTable{
		Cards:      "JJ",
		Percentage: 3,
	},
	PreflopTable{
		Cards:      "AQs",
		Percentage: 3.3,
	},
	PreflopTable{
		Cards:      "TT",
		Percentage: 3.8,
	},
	PreflopTable{
		Cards:      "AQo",
		Percentage: 4.7,
	},
	PreflopTable{
		Cards:      "99",
		Percentage: 5.1,
	},
	PreflopTable{
		Cards:      "AJs",
		Percentage: 5.4,
	},
	PreflopTable{
		Cards:      "88",
		Percentage: 5.9,
	},
	PreflopTable{
		Cards:      "ATs",
		Percentage: 6.2,
	},
	PreflopTable{
		Cards:      "AJo",
		Percentage: 7.1,
	},
	PreflopTable{
		Cards:      "77",
		Percentage: 7.5,
	},
	PreflopTable{
		Cards:      "66",
		Percentage: 8,
	},
	PreflopTable{
		Cards:      "ATo",
		Percentage: 8.9,
	},
	PreflopTable{
		Cards:      "A9s",
		Percentage: 9.2,
	},
	PreflopTable{
		Cards:      "55",
		Percentage: 9.7,
	},
	PreflopTable{
		Cards:      "A8s",
		Percentage: 10,
	},
	PreflopTable{
		Cards:      "KQs",
		Percentage: 10.3,
	},
	PreflopTable{
		Cards:      "44",
		Percentage: 10.7,
	},
	PreflopTable{
		Cards:      "A9o",
		Percentage: 11.6,
	},
	PreflopTable{
		Cards:      "A7s",
		Percentage: 11.9,
	},
	PreflopTable{
		Cards:      "KJs",
		Percentage: 12.2,
	},
	PreflopTable{
		Cards:      "A5s",
		Percentage: 12.5,
	},
	PreflopTable{
		Cards:      "A8o",
		Percentage: 13.4,
	},
	PreflopTable{
		Cards:      "A6s",
		Percentage: 13.7,
	},
	PreflopTable{
		Cards:      "A4s",
		Percentage: 14,
	},
	PreflopTable{
		Cards:      "33",
		Percentage: 14.5,
	},
	PreflopTable{
		Cards:      "KTs",
		Percentage: 14.8,
	},
	PreflopTable{
		Cards:      "A7o",
		Percentage: 15.7,
	},
	PreflopTable{
		Cards:      "A3s",
		Percentage: 16,
	},
	PreflopTable{
		Cards:      "KQo",
		Percentage: 16.9,
	},
	PreflopTable{
		Cards:      "A2s",
		Percentage: 17.2,
	},
	PreflopTable{
		Cards:      "A5o",
		Percentage: 18.1,
	},
	PreflopTable{
		Cards:      "A6o",
		Percentage: 19,
	},
	PreflopTable{
		Cards:      "A4o",
		Percentage: 19.9,
	},
	PreflopTable{
		Cards:      "KJo",
		Percentage: 20.8,
	},
	PreflopTable{
		Cards:      "QJs",
		Percentage: 21.1,
	},
	PreflopTable{
		Cards:      "A3o",
		Percentage: 22,
	},
	PreflopTable{
		Cards:      "22",
		Percentage: 22.5,
	},
	PreflopTable{
		Cards:      "K9s",
		Percentage: 22.8,
	},
	PreflopTable{
		Cards:      "A2o",
		Percentage: 23.7,
	},
	PreflopTable{
		Cards:      "KTo",
		Percentage: 24.6,
	},
	PreflopTable{
		Cards:      "QTs",
		Percentage: 24.9,
	},
	PreflopTable{
		Cards:      "K8s",
		Percentage: 25.2,
	},
	PreflopTable{
		Cards:      "K7s",
		Percentage: 25.5,
	},
	PreflopTable{
		Cards:      "JTs",
		Percentage: 25.8,
	},
	PreflopTable{
		Cards:      "K9o",
		Percentage: 26.7,
	},
	PreflopTable{
		Cards:      "K6s",
		Percentage: 27,
	},
	PreflopTable{
		Cards:      "QJo",
		Percentage: 27.9,
	},
	PreflopTable{
		Cards:      "Q9s",
		Percentage: 28.2,
	},
	PreflopTable{
		Cards:      "K5s",
		Percentage: 28.5,
	},
	PreflopTable{
		Cards:      "K8o",
		Percentage: 29.4,
	},
	PreflopTable{
		Cards:      "K4s",
		Percentage: 29.7,
	},
	PreflopTable{
		Cards:      "QTo",
		Percentage: 30.6,
	},
	PreflopTable{
		Cards:      "K7o",
		Percentage: 31.5,
	},
	PreflopTable{
		Cards:      "K3s",
		Percentage: 31.8,
	},
	PreflopTable{
		Cards:      "K2s",
		Percentage: 32.1,
	},
	PreflopTable{
		Cards:      "Q8s",
		Percentage: 32.4,
	},
	PreflopTable{
		Cards:      "K6o",
		Percentage: 33.3,
	},
	PreflopTable{
		Cards:      "J9s",
		Percentage: 33.6,
	},
	PreflopTable{
		Cards:      "K5o",
		Percentage: 34.5,
	},
	PreflopTable{
		Cards:      "Q9o",
		Percentage: 35.4,
	},
	PreflopTable{
		Cards:      "JTo",
		Percentage: 36.3,
	},
	PreflopTable{
		Cards:      "K4o",
		Percentage: 37.3,
	},
	PreflopTable{
		Cards:      "Q7s",
		Percentage: 37.6,
	},
	PreflopTable{
		Cards:      "T9s",
		Percentage: 37.9,
	},
	PreflopTable{
		Cards:      "Q6s",
		Percentage: 38.2,
	},
	PreflopTable{
		Cards:      "K3o",
		Percentage: 39.1,
	},
	PreflopTable{
		Cards:      "J8s",
		Percentage: 39.4,
	},
	PreflopTable{
		Cards:      "Q5s",
		Percentage: 39.7,
	},
	PreflopTable{
		Cards:      "K2o",
		Percentage: 40.6,
	},
	PreflopTable{
		Cards:      "Q8o",
		Percentage: 41.5,
	},
	PreflopTable{
		Cards:      "Q4s",
		Percentage: 41.8,
	},
	PreflopTable{
		Cards:      "J9o",
		Percentage: 42.7,
	},
	PreflopTable{
		Cards:      "Q3s",
		Percentage: 43,
	},
	PreflopTable{
		Cards:      "T8s",
		Percentage: 43.3,
	},
	PreflopTable{
		Cards:      "J7s",
		Percentage: 43.6,
	},
	PreflopTable{
		Cards:      "Q7o",
		Percentage: 44.5,
	},
	PreflopTable{
		Cards:      "Q2s",
		Percentage: 44.8,
	},
	PreflopTable{
		Cards:      "Q6o",
		Percentage: 45.7,
	},
	PreflopTable{
		Cards:      "98s",
		Percentage: 46,
	},
	PreflopTable{
		Cards:      "Q5o",
		Percentage: 46.9,
	},
	PreflopTable{
		Cards:      "J8o",
		Percentage: 47.8,
	},
	PreflopTable{
		Cards:      "T9o",
		Percentage: 48.7,
	},
	PreflopTable{
		Cards:      "J6s",
		Percentage: 49,
	},
	PreflopTable{
		Cards:      "T7s",
		Percentage: 49.3,
	},
	PreflopTable{
		Cards:      "J5s",
		Percentage: 49.6,
	},
	PreflopTable{
		Cards:      "Q4o",
		Percentage: 50.5,
	},
	PreflopTable{
		Cards:      "J4s",
		Percentage: 50.8,
	},
	PreflopTable{
		Cards:      "J7o",
		Percentage: 51.7,
	},
	PreflopTable{
		Cards:      "Q3o",
		Percentage: 52.6,
	},
	PreflopTable{
		Cards:      "97s",
		Percentage: 52.9,
	},
	PreflopTable{
		Cards:      "T8o",
		Percentage: 53.8,
	},
	PreflopTable{
		Cards:      "J3s",
		Percentage: 54.1,
	},
	PreflopTable{
		Cards:      "T6s",
		Percentage: 54.4,
	},
	PreflopTable{
		Cards:      "Q2o",
		Percentage: 55.4,
	},
	PreflopTable{
		Cards:      "J2s",
		Percentage: 55.7,
	},
	PreflopTable{
		Cards:      "87s",
		Percentage: 56,
	},
	PreflopTable{
		Cards:      "J6o",
		Percentage: 56.9,
	},
	PreflopTable{
		Cards:      "98o",
		Percentage: 57.8,
	},
	PreflopTable{
		Cards:      "T7o",
		Percentage: 58.7,
	},
	PreflopTable{
		Cards:      "96s",
		Percentage: 59,
	},
	PreflopTable{
		Cards:      "J5o",
		Percentage: 59.9,
	},
	PreflopTable{
		Cards:      "T5s",
		Percentage: 60.2,
	},
	PreflopTable{
		Cards:      "T4s",
		Percentage: 60.5,
	},
	PreflopTable{
		Cards:      "86s",
		Percentage: 60.8,
	},
	PreflopTable{
		Cards:      "J4o",
		Percentage: 61.7,
	},
	PreflopTable{
		Cards:      "T6o",
		Percentage: 62.6,
	},
	PreflopTable{
		Cards:      "97o",
		Percentage: 63.5,
	},
	PreflopTable{
		Cards:      "T3s",
		Percentage: 63.8,
	},
	PreflopTable{
		Cards:      "76s",
		Percentage: 64.1,
	},
	PreflopTable{
		Cards:      "95s",
		Percentage: 64.4,
	},
	PreflopTable{
		Cards:      "J3o",
		Percentage: 65.3,
	},
	PreflopTable{
		Cards:      "T2s",
		Percentage: 65.6,
	},
	PreflopTable{
		Cards:      "87o",
		Percentage: 66.5,
	},
	PreflopTable{
		Cards:      "85s",
		Percentage: 66.8,
	},
	PreflopTable{
		Cards:      "96o",
		Percentage: 67.7,
	},
	PreflopTable{
		Cards:      "T5o",
		Percentage: 68.6,
	},
	PreflopTable{
		Cards:      "J2o",
		Percentage: 69.5,
	},
	PreflopTable{
		Cards:      "75s",
		Percentage: 69.8,
	},
	PreflopTable{
		Cards:      "94s",
		Percentage: 70.1,
	},
	PreflopTable{
		Cards:      "T4o",
		Percentage: 71,
	},
	PreflopTable{
		Cards:      "65s",
		Percentage: 71.3,
	},
	PreflopTable{
		Cards:      "86o",
		Percentage: 72.2,
	},
	PreflopTable{
		Cards:      "93s",
		Percentage: 72.5,
	},
	PreflopTable{
		Cards:      "84s",
		Percentage: 72.9,
	},
	PreflopTable{
		Cards:      "95o",
		Percentage: 73.8,
	},
	PreflopTable{
		Cards:      "T3o",
		Percentage: 74.7,
	},
	PreflopTable{
		Cards:      "76o",
		Percentage: 75.6,
	},
	PreflopTable{
		Cards:      "92s",
		Percentage: 75.9,
	},
	PreflopTable{
		Cards:      "74s",
		Percentage: 76.2,
	},
	PreflopTable{
		Cards:      "54s",
		Percentage: 76.5,
	},
	PreflopTable{
		Cards:      "T2o",
		Percentage: 77.4,
	},
	PreflopTable{
		Cards:      "85o",
		Percentage: 78.3,
	},
	PreflopTable{
		Cards:      "64s",
		Percentage: 78.6,
	},
	PreflopTable{
		Cards:      "83s",
		Percentage: 78.9,
	},
	PreflopTable{
		Cards:      "94o",
		Percentage: 79.8,
	},
	PreflopTable{
		Cards:      "75o",
		Percentage: 80.7,
	},
	PreflopTable{
		Cards:      "82s",
		Percentage: 81,
	},
	PreflopTable{
		Cards:      "73s",
		Percentage: 81.3,
	},
	PreflopTable{
		Cards:      "93o",
		Percentage: 82.2,
	},
	PreflopTable{
		Cards:      "65o",
		Percentage: 83.1,
	},
	PreflopTable{
		Cards:      "53s",
		Percentage: 83.4,
	},
	PreflopTable{
		Cards:      "63s",
		Percentage: 83.7,
	},
	PreflopTable{
		Cards:      "84o",
		Percentage: 84.6,
	},
	PreflopTable{
		Cards:      "92o",
		Percentage: 85.5,
	},
	PreflopTable{
		Cards:      "43s",
		Percentage: 85.8,
	},
	PreflopTable{
		Cards:      "74o",
		Percentage: 86.7,
	},
	PreflopTable{
		Cards:      "72s",
		Percentage: 87,
	},
	PreflopTable{
		Cards:      "54o",
		Percentage: 87.9,
	},
	PreflopTable{
		Cards:      "64o",
		Percentage: 88.8,
	},
	PreflopTable{
		Cards:      "52s",
		Percentage: 89.1,
	},
	PreflopTable{
		Cards:      "62s",
		Percentage: 89.4,
	},
	PreflopTable{
		Cards:      "83o",
		Percentage: 90.3,
	},
	PreflopTable{
		Cards:      "42s",
		Percentage: 90.6,
	},
	PreflopTable{
		Cards:      "82o",
		Percentage: 91.6,
	},
	PreflopTable{
		Cards:      "73o",
		Percentage: 92.5,
	},
	PreflopTable{
		Cards:      "53o",
		Percentage: 93.4,
	},
	PreflopTable{
		Cards:      "63o",
		Percentage: 94.3,
	},
	PreflopTable{
		Cards:      "32s",
		Percentage: 94.6,
	},
	PreflopTable{
		Cards:      "43o",
		Percentage: 95.5,
	},
	PreflopTable{
		Cards:      "72o",
		Percentage: 96.4,
	},
	PreflopTable{
		Cards:      "52o",
		Percentage: 97.3,
	},
	PreflopTable{
		Cards:      "62o",
		Percentage: 98.2,
	},
	PreflopTable{
		Cards:      "42o",
		Percentage: 99.1,
	},
	PreflopTable{
		Cards:      "32o",
		Percentage: 100,
	},
}
