package modelos

import (
	"fmt"
	"time"
)

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

func (u *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	u.Id = id
	u.Name = name
	u.CreatedAt = createdAt
	u.Status = status
}

func (u *User) GetUser() {
	fmt.Println("Id.......: ", u.Id)
	fmt.Println("Name.....: ", u.Name)
	fmt.Println("CreatedAt: ", u.CreatedAt)
	fmt.Println("Status...: ", u.Status)
}
