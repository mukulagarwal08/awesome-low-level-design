package enums

type Status int

const (
	PENDING Status = iota + 1
	IN_PROGRESS
	COMPLETED
)

func (s Status) String() string {
	return [...]string{"PENDING", "IN-PROGRESS", "COMPLETED"}[s-1]
}
