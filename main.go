package main

import (
	"fmt"
	"gin_sample_web_app/infrastructure/dbclient"
	"time"
)

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	// fmt.Println(111)

	// testFuncStruct := modules.TestFunc{}
	// fmt.Println(context.TODO())
	// fmt.Print(testFuncStruct.TestRun(111))

	// r := config.DefineRoutes()
	// r.Run(":8080")

	// p := &operations.PostUserParams{
	// 	Body: &models.User{
	// 		ID:   1,
	// 		Name: "name_log",
	// 	},
	// }

	// p.SetTimeout(10 * time.Second)

	// transport := httptransport.New("localhost:8000", "api", nil)
	// client := apiclient.New(transport, strfmt.Default)

	// res, _ := client.Operations.PostUser(p)

	// log.Printf("&#v\n", res.Error())

	// vs := []modules.Stringfy{
	// 	&modules.Person{Name: "Taro", Age: 21},
	// 	&modules.Car{Number: "123345", Model: "111111"},
	// }

	// for _, v := range vs {
	// 	fmt.Println(v.ToString())
	// }

	// err := modules.RaiseError()
	// fmt.Println(err.Error())

	// e, ok := err.(*modules.MyError)
	// if ok {
	// 	fmt.Println(e.ErrCode)
	// }

	// s := []int{2, 3, 4, 5, 6}

	// fmt.Println(modules.Average(s))

	// u := new(modules.User)
	// u.Id = 1
	// u.Name = "test User"
	// u.Email = "example@example.com"
	// u.Created = time.Now()

	// bs, err := json.Marshal(u)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(bs))

	// fmt.Printf("%T\n", bs)

	// u2 := new(modules.User)

	// if err := json.Unmarshal(bs, &u2); err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(u2)

	// modules.PrintError(modules.DoError(modules.HOGEHOGE))
	// modules.PrintError(modules.DoError(modules.FUGAFUGA))
	// modules.PrintError(modules.DoError(modules.HOGEFUGA))
	// modules.CallPanic()
	// modules.PanicTest()
	// modules.RecoverPanic()
	// modules.RecoverPanicTestFunc()

	// c := modules.Car{
	// 	Number: "1111",
	// 	Model:  "11111111",
	// }

	// er := c.DoError()
	// switch e := er.(type) {
	// case modules.TestOverRideInterFace:
	// 	fmt.Print(1111)
	// 	fmt.Println(e)
	// default:
	// 	fmt.Println(11138)
	// }

	// fmt.Println(111)
	// err := modules.DoError(modules.HOGEHOGE)
	// // fmt.Println(err)
	// zap.S().Infof(
	// 	"test",
	// 	zap.Bool("success", false),
	// 	zap.String("method", "ttttt"),
	// 	zap.Error(err),
	// )
	// logger, _ := zap.NewDevelopment()
	// logger.Warn("Hello zap", zap.String("key", "value"), zap.Time("now", time.Now()))

	// v := modules.Vertex{1, 2}
	// fmt.Println(v.Abs())

	// var f modules.MyFloat
	// f = -0.3333
	// fmt.Println(f.Abs())
	// fmt.Println(-math.Sqrt2)

	// v := modules.Vertex{3, 4}
	// v.Scale(10.3)
	// fmt.Println(v)
	// fmt.Println(v.Abs())

	// f, err := os.Open("/tmp/hoge.txt")
	// if err != nil {
	// 	fmt.Println(111)
	// 	t := modules.TFuncHoge{Err: err}
	// 	switch e := t.Err.(type) {
	// 	case modules.TFunc:
	// 		fmt.Println(e)
	// 		fmt.Println("bbbbb")
	// 	default:
	// 		fmt.Println(e)
	// 		fmt.Println("aaaaaa")
	// 	}

	// }
	// fmt.Println(f)

	// m := modules.MyStringStruct{"bbbssaaa", "wwwbbbbbbp"}
	// fmt.Println(strings.Split(m.Y, "bbb"))
	// x := string(m.X)
	// fmt.Println(strings.Split(x, "s"))

	// _, err := os.Open("/tmp/hoge.txt")

	// e := modules.BookDatabase{Err: err}
	// _, ok := e.(modules.MyInterFace)
	// if ok {
	// 	fmt.Println(1)
	// }

	// var t modules.MyInterFace = &modules.BookDatabase{}
	// s := modules.Shop{}
	// tp := reflect.TypeOf(s)
	// fmt.Println(tp == reflect.TypeOf(modules.Shop{}))

	db := dbclient.ConnectGorm()
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&dbclient.User{})

	// user1 := User{Name: "テスト1太郎", Age: 20, Sex: "男"}
	// user2 := User{Name: "テスト二郎", Age: 25, Sex: "男"}
	// insertUsers := []User{user1, user2}
	// insert(insertUsers, db)
	user1 := dbclient.User{
		Name:         "test1",
		Age:          20,
		Sex:          "Man",
		MessageToken: "aaaaaaa",
	}
	fmt.Println(user1)
	time.Sleep(10 * time.Second)
	user, err := dbclient.FindLast(db)
	fmt.Println(user)
	fmt.Println(err)

}
