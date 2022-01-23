package main

import "fmt"

func main() {
	var game string
	fmt.Scanln(&game)

	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	res := 0
	for _, v := range results {
		if v == "w" {
			res += 3
		} else if v == "d" {
			res += 1
		} else if v == "l" {
			res += 0
		}
	}
	switch game {
	case "w":
		res += 3
	case "d":
		res += 1
	case "l":
		res += 0
	}
	fmt.Println(res)
}
