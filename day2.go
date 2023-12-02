package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs/day2.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	powerSum := 0
	var max = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	type Round struct {
		red   int
		green int
		blue  int
	}

	type Game struct {
		rounds []Round
	}

	var games []Game

	gameNumber := 1
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()

		splits := strings.Split(line, ":")
		rounds := strings.Split(splits[1], ";")
		var r []Round
		for _, round := range rounds {
			var red, green, blue = 0, 0, 0
			values := strings.Split(round, ",")
			for _, value := range values {
				entry := strings.Split(strings.TrimSpace(value), " ")
				if entry[1] == "red" {
					red, _ = strconv.Atoi(entry[0])
				} else if entry[1] == "green" {
					green, _ = strconv.Atoi(entry[0])
				} else if entry[1] == "blue" {
					blue, _ = strconv.Atoi(entry[0])
				}
			}
			r = append(r, Round{red, green, blue})
		}
		games = append(games, Game{r})

		localMax := Round{
			red: 0, green: 0, blue: 0,
		}

		valid := true
		for _, v := range r {
			if v.red > max["red"] || v.green > max["green"] || v.blue > max["blue"] {
				valid = false
				// break
			}
			if v.red > localMax.red {
				localMax.red = v.red
			}
			if v.green > localMax.green {
				localMax.green = v.green
			}
			if v.blue > localMax.blue {
				localMax.blue = v.blue
			}
		}
		if valid {
			total += gameNumber
		}

		powerSum += localMax.red * localMax.green * localMax.blue

		gameNumber += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Part 1:")
	log.Println(total)
	log.Println("Part 2:")
	log.Println(powerSum)
}
