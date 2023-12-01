package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Lets GOOOO")

	file, err := os.Open("inputs/day1.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		line := scanner.Text()

		// Part one
		// re := regexp.MustCompile(`\d+`)
		// digits := re.FindAllString(line, -1)
		// number, err := strconv.Atoi(digits[0][0:1] + digits[len(digits)-1][len(digits[len(digits)-1])-1:])

		// Part two
		replaced := strings.ReplaceAll(line, "one", "o1ne")
		replaced = strings.ReplaceAll(replaced, "two", "t2wo")
		replaced = strings.ReplaceAll(replaced, "three", "th3ree")
		replaced = strings.ReplaceAll(replaced, "four", "fo4ur")
		replaced = strings.ReplaceAll(replaced, "five", "fi5ve")
		replaced = strings.ReplaceAll(replaced, "six", "s6ix")
		replaced = strings.ReplaceAll(replaced, "seven", "se7ven")
		replaced = strings.ReplaceAll(replaced, "eight", "ei8ght")
		replaced = strings.ReplaceAll(replaced, "nine", "n9ine")

		re := regexp.MustCompile(`\d+`)
		digits := re.FindAllString(replaced, -1)
		number, err := strconv.Atoi(digits[0][0:1] + digits[len(digits)-1][len(digits[len(digits)-1])-1:])

		if err != nil {
			continue
		}
		total += number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(total)
}
