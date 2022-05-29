package modules

import "fmt"

const (
	HOGEHOGE = "ホゲホゲ"
	FUGAFUGA = "フガフガ"
	HOGEFUGA = "ホゲフガ"
)

type HogeError struct {
	Message string
}

func (e *HogeError) Error() string {
	return "This is Hoge Error"
}

type FugaError struct {
	Message string
	ErrCode int
}

func (e *FugaError) Error() string {
	return "This is Fuga Error"
}

func DoError(word string) error {
	if word == HOGEHOGE {
		return &HogeError{Message: word}
	}
	if word == FUGAFUGA {
		return &FugaError{Message: word, ErrCode: 123}
	}
	return nil
}

func PrintError(err error) {
	if err == nil {
		fmt.Println("no error")
		return
	}

	switch e := err.(type) {
	case *HogeError:
		fmt.Println("HogeError", err, "Message", e.Message)
	case *FugaError:
		fmt.Println("FugaError", err, "Message", e.Message, "ErrCode", e.ErrCode)
	default:
		fmt.Println("その他のエラー", err)
	}
}
