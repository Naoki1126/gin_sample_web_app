package main

import "gin_sample_web_app/sample"

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	sample.ErrorSample1()
	sample.ErrorSample2()
}
