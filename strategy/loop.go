package strategy

type Loop struct{}

func (l *Loop) Search(boxes []int, prisonerNumber int) bool {
	index := prisonerNumber
	guesses := len(boxes) / 2

	for i := 0; i < guesses; i++ {
		value := boxes[index]
		if value == prisonerNumber {
			return true
		}
		index = value
	}

	return false
}
