package main

import (
	"fmt"
	"gin_sample_web_app/infrastructure/dbclient"
	"gin_sample_web_app/sample"
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

	// db := dbclient.ConnectGorm()
	// defer db.Close()
	// db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&dbclient.User{})

	// // user1 := User{Name: "テスト1太郎", Age: 20, Sex: "男"}
	// // user2 := User{Name: "テスト二郎", Age: 25, Sex: "男"}
	// // insertUsers := []User{user1, user2}
	// // insert(insertUsers, db)
	// user1 := dbclient.User{
	// 	Name:         "test1",
	// 	Age:          20,
	// 	Sex:          "Man",
	// 	MessageToken: "aaaaaaa",
	// }
	// fmt.Println(user1)

	// user, err := dbclient.FindLast(db)
	// fmt.Println(user)
	// fmt.Println(err)

	// var i modules.TestOverRideInterFace

	// c := modules.Car{"sssss", "dddd"}
	// i = &c

	// t, ok := i.(*modules.Car)
	// fmt.Println(t)
	// if ok {
	// 	fmt.Println("ok")
	// }

	// now := time.Now()
	// tt := reflect.TypeOf(now)
	// fmt.Println(tt)

	db := dbclient.ConnectGorm()
	// u1 := dbclient.User{Name: "Taro", Age: 11}
	// u2 := dbclient.User{Name: "Taro", Age: 12}
	// u3 := dbclient.User{Name: "Taro", Age: 13}
	// dbclient.Insert(u1, db)
	// dbclient.Insert(u2, db)
	// dbclient.Insert(u3, db)

	// var users []dbclient.User
	// db.Find(&users).Debug()
	// for _, u := range users {
	// 	fmt.Println(u.Name)
	// }
	var result dbclient.MysqlSumResult
	var user dbclient.User
	fmt.Println(result.AgeSum)
	// r, err := db.Debug().Model(user).Select("sum(age) as AgeSum").()
	db.Model(user).Select("sum(age) as age_sum").Scan(&result)
	fmt.Println(result.AgeSum)
	// fmt.Println(r.Scan(&AgeSum))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(r)

	defer db.Close()

	profile := sample.MyProfile{FirstName: "Tanaka", LastName: "takashi", Gender: "male", Birthday: "19990909"}
	target := sample.TargetProfile{
		FirstName: "target",
		LastName:  "taro",
		Gender:    "male",
		Birthday:  "11111111",
	}
	var i1 interface{}
	var i2 interface{}
	i1 = profile
	i2 = target

	i3 := []sample.Human{&profile, &target}

	for _, u := range i3 {
		m := u.NickName()
		fmt.Println(m)
	}

	_, ok := i1.(sample.MyProfile)
	if ok {
		fmt.Println("success")
	}

	_, ok2 := i2.(sample.TargetProfile)

	if ok2 {
		fmt.Println("success")
	}

	err := profile.OpenFile()
	if err != nil {
		fmt.Println(err)
		// if errors.Is(err, &sample.MyProfile{}.Error()) {
		// 	fmt.Println(11111)
		// }
	}

	err2 := target.OpenFile()
	if err2 != nil {
		fmt.Println(err2)
	}

}
