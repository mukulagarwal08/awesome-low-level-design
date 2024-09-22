package enums

type Company int

const (
	MARUTI = iota + 1
	HYUNDAI
	MAHINDRA
)

func (c Company) String() string {
	return [...]string{"Maruti", "Hyundai", "Mahindra"}[c-1]
}
