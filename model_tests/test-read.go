package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hpcloud/hsm-cli/Godeps/_workspace/src/github.com/davecgh/go-spew/spew"
	"github.com/dagnello/go-swagger-test/model_tests/models_hsm/models"
)

func main() {
	// read original json file
	dat, err := ioutil.ReadFile("test_input.json")
	if err != nil {
		fmt.Println("err1: " + err.Error())
		return
	}

	myService := models.Service{}


	spew.Dump(dat)

	// unmarshal json to models.Service object
	err = json.Unmarshal(dat, &myService)
	if err != nil {
		fmt.Println("err2: " + err.Error())
		return
	}

	spew.Dump(myService.Parameters)

	// marshal models.Service object to json and write to file
	b, err := json.Marshal(myService)
	if err != nil {
		fmt.Println("err3: " + err.Error())
		return
	}

	f, err := os.Create("test_output.json")
	if err != nil {
		fmt.Println("err4: " + err.Error())
		return
	}
	defer f.Close()
	f.Write(b)

	// read marshalled output from test_output.json file
	dat2, err := ioutil.ReadFile("test_output.json")
	if err != nil {
		fmt.Println("err5: " + err.Error())
		return
	}

	myService2 := models.Service{}

	// unmarshal json to models.Service object (fails because 'type' field is not in generator struct - from dat2)
	err = json.Unmarshal(dat2, &myService2)
	if err != nil {
		fmt.Println("err6: " + err.Error())
		return
	}
	spew.Dump(myService2.Parameters)
}