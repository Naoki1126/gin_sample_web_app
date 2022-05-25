package main

import (
	"gin_sample_web_app/swagger/client/operations"
	"gin_sample_web_app/swagger/models"
	"log"
	"time"

	apiclient "gin_sample_web_app/swagger/client"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func main() {
	// fmt.Println(111)

	// testFuncStruct := modules.TestFunc{}
	// fmt.Println(context.TODO())
	// fmt.Print(testFuncStruct.TestRun(111))

	// r := config.DefineRoutes()
	// r.Run(":8080")

	p := &operations.PostUserParams{
		Body: &models.User{
			ID:   1,
			Name: "name_log",
		},
	}

	p.SetTimeout(10 * time.Second)

	transport := httptransport.New("localhost:8000", "api", nil)
	client := apiclient.New(transport, strfmt.Default)

	res, _ := client.Operations.PostUser(p)

	log.Printf("&#v\n", res.Error())
}
