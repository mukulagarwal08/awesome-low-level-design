package enitities

type UserInventory struct {
	UserData map[int]*User
}

func NewUserInventory() *UserInventory {
	return &UserInventory{
		UserData: make(map[int]*User),
	}
}

func (u *UserInventory) GetUserById(id int) *User {
	return u.UserData[id]
}
