package strategy

type AlwaysFalse struct{}

func (l *AlwaysFalse) Search(_boxes []int, _prisonerNumber int) bool {
	return false
}
