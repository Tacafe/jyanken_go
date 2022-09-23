package jyanken

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetComputerHand() string {
	rand.Seed(time.Now().UnixNano())
	random_num := rand.Intn(4)
	return jyankenHandMap()[random_num-1]
}

func GetPlayerHand() string {
	fmt.Printf("\t勝負する手を選択してください\n")

	promptPlayerHand()

	var player_input string
	reader := bufio.NewReader(os.Stdin)
	player_input, _ = reader.ReadString('\n')
	player_input = strings.TrimSuffix(player_input, "\n")

	for IsValid(player_input) == false {
		promptPlayerHand()
		player_input, _ = reader.ReadString('\n')
		player_input = strings.TrimSuffix(player_input, "\n")
	}

	player_choice, _ := strconv.Atoi(player_input)
	return jyankenHandMap()[player_choice-1]
}

func jyankenHandNames() [3]string {
	return [3]string{"グー", "チョキ", "パー"}
}

func jyankenHandMap() map[int]string {
	m := make(map[int]string)
	for i, hand_name := range jyankenHandNames() {
		m[i] = hand_name
	}
	return m
}

func getStoredKeys(m map[int]string) []int {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func promptPlayerHand() {
	var prmpt string
	m := jyankenHandMap()
	for id := range getStoredKeys(m) {
		item := m[id]
		prmpt += fmt.Sprintf("\t%s\t%d を入力\n", item, id+1)
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
