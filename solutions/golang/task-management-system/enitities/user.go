package enitities

type User struct {
	id    int
	Email string
	Name  string
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
}
