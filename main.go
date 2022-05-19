package main

import (
	"fmt"
	"strings"
	"example.com/packages/words"
)



func main() {
	// pick random word from the function
	fmt.Println("rules:")
	fmt.Println("enter word that contains of 5 letters and you get different answers based on each letter:")
	fmt.Println("O => when the character at the right place in the word")
	fmt.Println("X => when the character is not contained in the word")
	fmt.Println("P => when the character is contained in the word but not at the right place")
	goal_word := words.PickWord()
	game_over := 0
	var word_input string
	for game_over < 5 {
		fmt.Println("enter a word")
		fmt.Scanln(&word_input)
		// check if the input has more than 5 letters 
		if len(word_input) != 5 {
			fmt.Println("enter word of 5 letters only")
			continue
		}
		// check if you input invalid word
		if words.CheckWord(word_input) == false {
			fmt.Println("enter correct word")
			continue
		}
		if word_input == goal_word {
			fmt.Println(word_input)
			break;
		}
		var x []string
		for i := 0; i < len(word_input); i++ {
			if word_input[i:i+1] == goal_word[i:i+1] {
				x = append(x, "O");
			} else if strings.Contains(goal_word, word_input[i:i+1]) == true {
				x = append(x, "P");
			} else {
				x = append(x, "X");
			}
		}
		s := strings.Split(word_input, "")
		fmt.Println(x)
		fmt.Println(s)
		game_over += 1
	}
	if game_over >= 5{
		fmt.Println("you lost and the word was")
	} else {
		fmt.Println("you won and the word was")
	}
	fmt.Println(goal_word)
	


}