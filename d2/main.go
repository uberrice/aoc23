package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	id int
	c  []cubes
}

type cubes struct {
	r int
	b int
	g int
}

func process_games(text string) []game {
	var games []game
	// var game_id int = 0

	for _, line := range strings.Split(text, "\r\n") {
		var g game
		var line_split []string = strings.Split(line, ":")
		fmt.Println(line_split[0])
		g.id, _ = strconv.Atoi(strings.Split(line_split[0], " ")[1])
		fmt.Sscanf(line_split[0], "Game %d", g.id)
		var occ []string = strings.Split(line_split[1], ";")

		for _, occurence := range occ {
			var occ_split []string = strings.Split(occurence, ",")
			var c cubes

			for _, cube := range occ_split {
				color_set := strings.Split(cube, " ")
				switch color_set[2] {
				case "red":
					c.r, _ = strconv.Atoi(color_set[1])
				case "blue":
					c.b, _ = strconv.Atoi(color_set[1])
				case "green":
					c.g, _ = strconv.Atoi(color_set[1])
				}
			}
			g.c = append(g.c, c)
		}
		games = append(games, g)
	}
	return games
}

func check_game(g game, check cubes) int {
	for _, c := range g.c {
		if c.r > check.r || c.b > check.b || c.g > check.g {
			fmt.Printf("game %d not ok, r: %d %d g: %d %d b: %d %d\n", g.id, c.r, check.r, c.g, check.g, c.b, check.b)
			return 0
		}
	}
	return g.id
}

func check_games(games []game, check cubes) []int {
	var res []int
	for _, g := range games {
		var game_id int = check_game(g, check)
		if game_id != 0 {
			res = append(res, game_id)
		}
	}
	return res
}

func game_max_cubes(g game) cubes {
	var max_c cubes
	for _, c := range g.c {
		if c.r > max_c.r {
			max_c.r = c.r
		}
		if c.b > max_c.b {
			max_c.b = c.b
		}
		if c.g > max_c.g {
			max_c.g = c.g
		}
	}
	return max_c
}

func cube_power(c cubes) int {
	return c.r * c.b * c.g
}

func main() {
	var fc []byte
	fc, err := os.ReadFile("d2_in.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var input = string(fc)

	var games = process_games(input)
	fmt.Printf("games: %v\n", games)
	var check cubes = cubes{r: 12, b: 14, g: 13}
	var res []int = check_games(games, check)
	fmt.Printf("res: %v\n", res)
	var res_sum int = 0
	for _, r := range res {
		res_sum += r
	}
	fmt.Printf("res_sum: %d\n", res_sum)

	var sum_cube_power int = 0
	for _, g := range games {
		var max_c cubes = game_max_cubes(g)
		sum_cube_power += cube_power(max_c)
	}
	fmt.Printf("sum_cube_power: %d\n", sum_cube_power)
}
