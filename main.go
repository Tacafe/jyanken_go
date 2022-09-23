package main

import (
	"fmt"

	"jyanken_go/jyanken"
)

func main() {
	initialMatch()
}

func initialMatch() {
	fmt.Printf("\nさいしょは\t✊\n")
	player_hand := jyanken.GetPlayerHand()
	com_hand := jyanken.GetComputerHand()

	fmt.Printf("じゃーんけーん....ポンっ!!!\n")
	fmt.Printf("YOU:\t%s\n", player_hand)
	fmt.Printf("COM:\t%s\n\n", com_hand)
}

	fmt.Printf("さいしょは✊\n")
	fmt.Printf("じゃーんけーん: %s\n", player_hand)
}
