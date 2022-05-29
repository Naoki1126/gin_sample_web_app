package modules

import (
	"encoding/json"
	"fmt"
	"time"
)

type A struct{}

type User struct {
	Id      int       `json:"id, string"`
	Name    string    `json:"name"`
	Email   string    `json: "email"`
	Created time.Time `json: "created"`
	A       A         `json:A`
}

func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + u.Name,
	})

	return v, err
}

func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}

	var u2 User2

	err := json.Unmarshal(b, &u2)

	if err != nil {
		fmt.Println(err)
	}

	u.Name = u2.Name + "!"

	return err
}
