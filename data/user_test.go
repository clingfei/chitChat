package data

import (
	"fmt"
	"testing"
)

var users = User{
	Name: "alicse",
	Email: "12345678@qq.com",
	Password: "123456",
}
func setup() {
	UserDeleteAll()
}

func TestUser_Create(t *testing.T) {
	setup()
	if err := users.Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	fmt.Println(users.Uuid)
	if users.Id == 0 {
		t.Errorf("No id or created_at in user")
	}
	u, err := UserByEmail(users.Email)
	if err != nil {
		t.Error(err, "User not created.")
	}
	if users.Email != u.Email {
		t.Errorf("User retrieved is not the same as the one created.")
	}
}