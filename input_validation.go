package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// validates a given value for option
// option_to_validate: human-readable value shown to the user to prompt for validation
func CorrectStringValidate(option_to_validate string) string {
	validate := true
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	var choice string
	scanner.Scan()
	choice = scanner.Text()
	for validate {
		fmt.Printf("Are you sure you want '%s' for '%s'\n[Y]Yes, [N]No: ", choice, option_to_validate)
		scanner.Scan()
		input = scanner.Text()
		input = strings.ToUpper(input)
		if input == "Y" {
			validate = false
		} else if input == "N" {
			fmt.Printf("Please enter a new value for %s: ", option_to_validate)
			scanner.Scan()
			choice = scanner.Text()
		} else {
			fmt.Println("Invalid key entered.")
		}
	}
	return choice
}

/*
Prompts gives the user yes or no question based on a given message
returns true if yes, false if no
*/
func YesNo(message string) bool {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s\n[Y]Yes [N]No: ", message)
		scanner.Scan()
		input = scanner.Text()
		input = strings.ToUpper(input)
		if input == "Y" {
			return true
		} else if input == "N" {
			return false
		} else {
			fmt.Println("Invalid key entered")
		}
	}
}

/*
Prompts the user to given an integer value
Can limit range by adding min or max (geq or leq)
Args: min, max
*/
func GetInt(message string, rng ...int) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Please enter a integer value for %s\n", message)
		scanner.Scan()
		input := scanner.Text()
		val, err := strconv.Atoi(input)
		if err == nil {
			if len(rng) == 1 {
				if val >= rng[0] {
					return val
				} else {
					fmt.Printf("Please enter a value that is greater than %d\n", rng[0])
				}
			} else if len(rng) == 2 {
				if val >= rng[0] && val <= rng[1] {
					return val
				} else {
					fmt.Printf("Please enter a value that is greater than %d and less than %d\n", rng[0], rng[1])
				}
			} else {
				confirm := fmt.Sprintf("Are you sure you would like %d for %s", val, message)
				if YesNo(confirm) {
					return val
				}
			}
		} else {
			fmt.Printf("Please enter an integer value less than %d", rng)
		}
	}
}
