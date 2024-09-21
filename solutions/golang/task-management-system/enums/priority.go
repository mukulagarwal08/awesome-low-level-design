package enums

type Priority int

const (
	SEV_ZERO Priority = iota + 1
	SEV_ONE
	SEV_TWO
)

func (p Priority) String() string {
	return [...]string{"SEV_ZERO", "SEV_ONE", "SEV_TWO"}[p]
}
