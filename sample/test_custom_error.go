package sample

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func ErrorSample1() {
	if _, err := os.Open("nos-existing"); err != nil {
		var pathError *fs.PathError

		wrapedErr := fmt.Errorf("err is %w", err)
		if errors.As(wrapedErr, &pathError) {
			fmt.Println("errors.As(): Failed at Path:", pathError.Path)
		} else {
			fmt.Println(err)
		}

		if errors.Is(wrapedErr, pathError) {
			fmt.Println("errors.is():Failed at path: ", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}

type SampleError2 struct {
	message string
	err     error
}

func (se2 *SampleError2) Error() string {
	return se2.message
}

func (se2 *SampleError2) Unwrap() error { // 追加
	return se2.err
}

func ErrorSample2() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError

		wrapedErr := &SampleError2{message: "this is wraped err", err: err}
		if errors.As(wrapedErr, &pathError) {
			fmt.Println("erros.As(): Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}

		if errors.Is(wrapedErr, pathError) {
			fmt.Println("errors.Is(): Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}
