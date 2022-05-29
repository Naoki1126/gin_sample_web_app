package modules

import "fmt"

func CallPanic() {
	fmt.Println("Start Panic function")
	panic("Panic Panic")
	fmt.Println("End Panic Function")
}

func PanicTest() {
	text := "hoge"
	fmt.Println("1:", text[0:2])
	// fmt.Println("2:", text[5:])
	fmt.Println("end")

	var p *int
	fmt.Printf("%p\n", p)
	// fmt.Printf("%d\n", *p)

	var x interface{} = "hogehogehoge"
	i, ok := x.(int)

	if ok {
		fmt.Println(i)
	}
}

func RecoverPanic() {
	fmt.Println("recover Panic start")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover", err)
		}
	}()
	panic("I am Panic !!!!!!")
	fmt.Println("after panic")
}

func RecoverPanicTestFunc() {
	fmt.Println("recover Panic test func")

	var x interface{} = "testetetetet"
	i := x.(int)
	fmt.Println(i)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(1111111111)
			fmt.Println("err : ", err)
		}
	}()

}
