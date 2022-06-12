package sample

import (
	"fmt"
	"os"
)

type Human interface {
	NickName() string
	IsMale() bool
	OpenFile() error
	// error
}

type MyProfile struct {
	FirstName string
	LastName  string
	Gender    string
	Birthday  string
}

func (u *MyProfile) OpenFile() error {
	_, err := os.Open("examle.txt")
	if err != nil {
		return err
	}
	return nil
}

func (u *MyProfile) NickName() string {
	return u.FirstName[:1] + u.LastName[:1]
}

func (u *MyProfile) IsMale() bool {
	return u.Gender == "male"
}

func (u *MyProfile) Error() error {
	return fmt.Errorf("this is myprofile error")
}

type TargetProfile struct {
	FirstName string
	LastName  string
	Gender    string
	Birthday  string
}

func (t *TargetProfile) NickName() string {
	return t.FirstName + t.LastName
}

func (t *TargetProfile) IsMale() bool {
	return t.Gender == "male"
}

func (t *TargetProfile) Error() error {
	return fmt.Errorf("This is TargetProfile erorr")
}

func (t *TargetProfile) OpenFile() error {
	_, err := os.Open("examle.txt")
	if err != nil {
		return err
	}
	return nil
}
