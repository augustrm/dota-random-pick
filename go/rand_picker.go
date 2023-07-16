package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

var heroes = [124]string{
	"Abaddon",
	"Alchemist",
	"Ancient Apparition",
	"Anti-Mage",
	"Arc Warden",
	"Axe",
	"Bane",
	"Batrider",
	"Beastmaster",
	"Bloodseeker",
	"Bounty Hunter",
	"Brewmaster",
	"Bristleback",
	"Broodmother",
	"Centaur Warrunner",
	"Chaos Knight",
	"Chen",
	"Clinkz",
	"Clockwerk",
	"Crystal Maiden",
	"Dark Seer",
	"Dark Willow",
	"Dawnbreaker",
	"Dazzle",
	"Death Prophet",
	"Disruptor",
	"Doom",
	"Dragon Knight",
	"Drow Ranger",
	"Earth Spirit",
	"Earthshaker",
	"Elder Titan",
	"Ember Spirit",
	"Enchantress",
	"Enigma",
	"Faceless Void",
	"Grimstroke",
	"Gyrocopter",
	"Hoodwink",
	"Huskar",
	"Invoker",
	"Io",
	"Jakiro",
	"Juggernaut",
	"Keeper of the Light",
	"Kunkka",
	"Legion Commander",
	"Leshrac",
	"Lich",
	"Lifestealer",
	"Lina",
	"Lion",
	"Lone Druid",
	"Luna",
	"Lycan",
	"Magnus",
	"Marci",
	"Mars",
	"Medusa",
	"Meepo",
	"Mirana",
	"Monkey King",
	"Morphling",
	"Muerta",
	"Naga Siren",
	"Nature's Prophet",
	"Necrophos",
	"Night Stalker",
	"Nyx Assassin",
	"Ogre Magi",
	"Omniknight",
	"Oracle",
	"Outworld Destroyer",
	"Pangolier",
	"Phantom Assassin",
	"Phantom Lancer",
	"Phoenix",
	"Primal Beast",
	"Puck",
	"Pudge",
	"Pugna",
	"Queen of Pain",
	"Razor",
	"Riki",
	"Rubick",
	"Sand King",
	"Shadow Demon",
	"Shadow Fiend",
	"Shadow Shaman",
	"Silencer",
	"Skywrath Mage",
	"Slardar",
	"Slark",
	"Snapfire",
	"Sniper",
	"Spectre",
	"Spirit Breaker",
	"Storm Spirit",
	"Sven",
	"Techies",
	"Templar Assassin",
	"Terrorblade",
	"Tidehunter",
	"Timbersaw",
	"Tinker",
	"Tiny",
	"Treant Protector",
	"Troll Warlord",
	"Tusk",
	"Underlord",
	"Undying",
	"Ursa",
	"Vengeful Spirit",
	"Venomancer",
	"Viper",
	"Visage",
	"Void Spirit",
	"Warlock",
	"Weaver",
	"Windranger",
	"Winter Wyvern",
	"Witch Doctor",
	"Wraith King",
	"Zeus",
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func cantor_pair(x int, y int) int {
	return (((x + y) * (x + y + 1)) / 2) + y
}

func instance_cantor_index() []int {
	cantor_index := []int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 31; j++ {
			cantor_index = append(cantor_index, cantor_pair(i, j))
		}
	}
	sort.Ints(cantor_index)
	return cantor_index
}

func get_adjacent_mn(m int, n int) [][]int {
	/*(i % n) + n) % n
	((n - 1) % 31) + 31) % 31*/
	top := []int{(((m - 1) % 4) + 4) % 4, n}
	bottom := []int{(((m + 1) % 4) + 4) % 4, n}
	left := []int{m, (((n - 1) % 31) + 31) % 31}
	right := []int{m, (((n + 1) % 31) + 31) % 31}
	adj_nodes := [][]int{top, bottom, left, right}
	return adj_nodes
}

type Pick struct {
	hero       string
	alt_heroes [4]string
}

func do_picks(k int) []Pick {
	Picks := []Pick{}
	cantor_index := instance_cantor_index()
	for i := 0; i < k; i++ {
		m := rand.Intn(4)
		n := rand.Intn(31)

		rand_pair := [2]int{m, n}
		adj_nodes := get_adjacent_mn(rand_pair[0], rand_pair[1])

		cantor_rand_pick := cantor_pair(rand_pair[0], rand_pair[1])
		cantor_adj_nodes := []int{
			cantor_pair(adj_nodes[0][0], adj_nodes[0][1]),
			cantor_pair(adj_nodes[1][0], adj_nodes[1][1]),
			cantor_pair(adj_nodes[2][0], adj_nodes[2][1]),
			cantor_pair(adj_nodes[3][0], adj_nodes[3][1]),
		}

		hero_index := indexOf(cantor_rand_pick, cantor_index)
		hero_adj_index := []int{
			indexOf(cantor_adj_nodes[0], cantor_index),
			indexOf(cantor_adj_nodes[1], cantor_index),
			indexOf(cantor_adj_nodes[2], cantor_index),
			indexOf(cantor_adj_nodes[3], cantor_index),
		}

		hero := heroes[hero_index]
		adj_heroes := [4]string{
			heroes[hero_adj_index[0]],
			heroes[hero_adj_index[1]],
			heroes[hero_adj_index[2]],
			heroes[hero_adj_index[3]],
		}
		var P Pick
		P.hero = hero
		P.alt_heroes = adj_heroes
		Picks = append(Picks, P)
	}
	return Picks
}

func main() {
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	picks := do_picks(k)

	for i := 0; i < k; i++ {
		fmt.Println("Hero: ", picks[i].hero, "\nAlt Heroes: ", picks[i].alt_heroes[0], ", ", picks[i].alt_heroes[1], ", ", picks[i].alt_heroes[2], ", ", picks[i].alt_heroes[3], "\n ")
	}
}
