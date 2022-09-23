package jyanken

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetPlayerHand() string {
	fmt.Printf("勝負する手を選択してください\n")

	PromptJyankenImput()

	var player_input string
	reader := bufio.NewReader(os.Stdin)
	player_input, _ = reader.ReadString('\n')
	player_input = strings.TrimSuffix(player_input, "\n")

	for IsValid(player_input) == false {
		PromptJyankenImput()
		player_input, _ = reader.ReadString('\n')
		player_input = strings.TrimSuffix(player_input, "\n")
	}

	player_choice, _ := strconv.Atoi(player_input)
	return JyankenHandMap()[player_choice-1]
}

func jyanken_hand_names() [3]string {
	return [3]string{"グー", "チョキ", "パー"}
}

func JyankenHandMap() map[int]string {
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

func PromptJyankenImput() {
	var prmpt string
	m := JyankenHandMap()
	for id := range get_sorted_keys(m) {
		item := m[id]
		prmpt += fmt.Sprintf("%s\t%d を入力\n", item, id+1)
	}
	fmt.Print(prmpt)
}

func IsValid(player_input string) bool {
	i, err := strconv.Atoi(player_input)
	if err != nil || i < 1 || 3 < i {
		return false
	}
	return true
}
