package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	var result string
	var initialPosition []int
	currentPosition := []int{4, 1}
	label := "Select Day"
	pattern := [][]string{
		{"#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", ".", ".", ".", ".", "$", ".", "#"},
		{"#", ".", "#", "#", "#", ".", "$", "#"},
		{"#", ".", ".", ".", "#", ".", "#", "#"},
		{"#", "x", "#", ".", "$", ".", "$", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#"},
	}
	// number of rows in s2
	n := len(pattern)

	// number of columns in s2
	m := len(pattern[0])
	for {
		fmt.Println("your current position: ")
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Print(pattern[i][j] + " ")
			}
			fmt.Print("\n")
		}

		if result != "exit" {
			initialPosition = currentPosition
			pattern[initialPosition[0]][initialPosition[1]] = "."
			option := []string{}
			if pattern[initialPosition[0]-1][initialPosition[1]] != "#" {
				option = append(option, "up")
			}
			if pattern[initialPosition[0]][initialPosition[1]+1] != "#" {
				option = append(option, "right")
			}
			if pattern[initialPosition[0]+1][initialPosition[1]] != "#" {
				option = append(option, "down")
			}
			option = append(option, "exit")
			test, err := prompt(label, option)
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
			}
			if test == "" {
				break
			}
			result = test

		}

		if result == "up" {
			currentPosition[0] -= 1
		}

		if result == "right" {
			currentPosition[1] += 1
		}

		if result == "exit" {
			break
		}

		if result == "down" {
			currentPosition[0] += 1
		}

		if currentPosition[0] == 2 && currentPosition[1] == 6 {
			pattern[currentPosition[0]][currentPosition[1]] = "W"
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					fmt.Print(pattern[i][j] + " ")
				}
				fmt.Print("\n")
			}
			fmt.Println("CONGRATULATION YOU HAVE FOUND THE TREASURE")
			break
		}

		if pattern[currentPosition[0]][currentPosition[1]] == "$" {
			label = "nice try buddy keep up the good work, sorry this is not the treasure, move forward"
		}
		pattern[currentPosition[0]][currentPosition[1]] = "x"
	}
	// fmt.Printf("Prompt result %v\n", result)
}

func prompt(label string, test []string) (result string, err error) {
	prompt := promptui.Select{
		Label: label,
		Items: test,
	}

	_, result, err1 := prompt.Run()

	if err1 != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return result, err
	}
	return result, err
}
