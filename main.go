package main

import (
	"fmt"

	"jyanken_go/jyanken"
)

func main() {
	var winner string
	winner = initialMatch()

	for winner == "" {
		winner = aikoMatch()
	}

	fmt.Printf("%sの勝ちです。\n", winner)
}

func initialMatch() string {
	fmt.Printf("\nさいしょは\t✊\n")
	player_hand := jyanken.GetPlayerHand()
	com_hand := jyanken.GetComputerHand()

	fmt.Printf("じゃーんけーん....ポンっ!!!\n")
	fmt.Printf("YOU:\t%s\n", player_hand)
	fmt.Printf("COM:\t%s\n\n", com_hand)
	return getResult(player_hand, com_hand)
}

func aikoMatch() string {
	fmt.Printf("あーいこーで....\n")
	player_hand := jyanken.GetPlayerHand()
	com_hand := jyanken.GetComputerHand()
	fmt.Printf("...しょっ!!!\n")
	fmt.Printf("YOU:\t%s\n", player_hand)
	fmt.Printf("COM:\t%s\n\n", com_hand)
	return getResult(player_hand, com_hand)
}

func getResult(player_hand string, com_hand string) string {
	switch player_hand {
	case "グー":
		switch com_hand {
		case "グー":
			return ""
		case "チョキ":
			return "YOU"
		case "パー":
			return "COM"
		}

	case "チョキ":
		switch com_hand {
		case "グー":
			return "COM"
		case "チョキ":
			return ""
		case "パー":
			return "YOU"
		}

	case "パー":
		switch com_hand {
		case "グー":
			return "YOU"
		case "チョキ":
			return "COM"
		case "パー":
			return ""
		}
	}
	return ""
}
