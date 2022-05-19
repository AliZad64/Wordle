package words

import (
	"math/rand"
	"time"
)
func PickWord() string {
	rand.Seed(time.Now().Unix())
	rand_index := rand.Intn(len(all_words))
	rand_word := all_words[rand_index]
	return rand_word
}

func CheckWord(word string) bool {
	for _, b := range all_words {
        if b == word {
            return true
        }
    }
    return false
}