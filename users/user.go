package users

import (
	"time"

	"github.com/maalpu/godesde0/modelos"
)

func Alta(id int, name string) {
	u := new(modelos.User)
	u.AddUser(id, name, time.Now(), true)
	u.GetUser()
}
