package strategy

type AlwaysTrue struct{}

func (l *AlwaysTrue) Search(_boxes []int, _prisonerNumber int) bool {
	return true
}
