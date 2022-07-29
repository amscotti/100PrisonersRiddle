package strategy

import (
	"math/rand"
)

type Random struct{}

func (l *Random) Search(boxes []int, prisonerNumber int) bool {
	index := prisonerNumber
	guessCount := len(boxes) / 2
	guesses := rand.Perm(len(boxes))

	for i := 0; i < guessCount; i++ {
		value := boxes[guesses[index]]
		if value == prisonerNumber {
			return true
		}
	}

	return false
}
