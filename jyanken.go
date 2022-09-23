package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var player_hand string
	player_hand = get_player_hand()

	fmt.Printf("さいしょは✊\n")
	fmt.Printf("じゃーんけーん: %s\n", player_hand)
}

func get_player_hand() string {
	fmt.Printf("勝負する手を選択してください\n")

	prompt_input_jyanken_hand()

	var player_input string
	reader := bufio.NewReader(os.Stdin)
	player_input, _ = reader.ReadString('\n')
	player_input = strings.TrimSuffix(player_input, "\n")

	for is_valid_input(player_input) == false {
		prompt_input_jyanken_hand()
		player_input, _ = reader.ReadString('\n')
		player_input = strings.TrimSuffix(player_input, "\n")
	}

	player_choice, _ := strconv.Atoi(player_input)
	return jyanken_hand_map()[player_choice-1]
}

func jyanken_hand_names() [3]string {
	return [3]string{"グー", "チョキ", "パー"}
}

func jyanken_hand_map() map[int]string {
	m := make(map[int]string)
	for i, hand_name := range jyanken_hand_names() {
		m[i] = hand_name
	}
	return m
}

func get_sorted_keys(m map[int]string) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func prompt_input_jyanken_hand() {
	var prmpt string
	m := jyanken_hand_map()
	for id := range get_sorted_keys(m) {
		item := m[id]
		prmpt += fmt.Sprintf("%s\t%d を入力\n", item, id+1)
	}
	fmt.Print(prmpt)
}

func is_valid_input(player_input string) bool {
	i, err := strconv.Atoi(player_input)
	if err != nil || i < 1 || 3 < i {
		return false
	}
	return true
}
