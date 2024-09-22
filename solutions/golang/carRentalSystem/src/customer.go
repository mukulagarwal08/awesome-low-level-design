package src

type Customer struct {
	Name    string
	Phone   int64
	license string
}

func NewCustomer(name string, phone int64, license string) *Customer {
	return &Customer{
		Name:    name,
		Phone:   phone,
		license: license,
	}
}
