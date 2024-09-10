package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello! I'm your weather assistant.")
	fmt.Println("I'll help you decide if you need a raincoat today.")

	reader := bufio.NewReader(os.Stdin)

	var chanceOfRain int
	for {
		fmt.Print("What's the chance of rain today? (0-100%): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		input = strings.TrimSpace(input)
		chance, err := strconv.Atoi(input)
		if err == nil && chance >= 0 && chance <= 100 {
			chanceOfRain = chance
			break
		}
		fmt.Println("Please enter a valid number between 0 and 100.")
	}

	fmt.Print("Is it currently raining? (yes/no): ")
	currentCondition, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	currentCondition = strings.TrimSpace(strings.ToLower(currentCondition))

	if chanceOfRain > 50 || currentCondition == "yes" {
		fmt.Println("It's likely to rain or already raining. You should wear a raincoat!")
	} else if chanceOfRain >= 30 && chanceOfRain <= 50 {
		fmt.Println("There's a moderate chance of rain. It might be a good idea to bring a raincoat just in case.")
	} else {
		fmt.Println("It's unlikely to rain. You probably don't need a raincoat today.")
	}

	fmt.Println("Stay dry and have a great day!")
}
