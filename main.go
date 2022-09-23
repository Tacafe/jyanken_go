package main

import (
	"fmt"

	"jyanken_go/jyanken"
)

func main() {
	var player_hand string
	player_hand = jyanken.GetPlayerHand()

	fmt.Printf("さいしょは✊\n")
	fmt.Printf("じゃーんけーん: %s\n", player_hand)
}
