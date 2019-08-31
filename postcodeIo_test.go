package main

import (
	"fmt"
	"github.com/beemi/postcode-io-tests-golang/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestPostCodeLatLong(t *testing.T) {
	url := config.PostCodeIOEndPoint() + "/postcodes/RM17%206EY"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		assert.Fail(t, "Error while client sending request")
	}
	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode, "Get Postcode lat long Api failed")
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	// JSON schema validation
	//schemaLoader := gojsonschema.NewReferenceLoader("file:///Users/rajbeemi/projects/personal/postcode-io-tests-golang/schemas/postcode_lat_long.json")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + dir + "/schemas/postcode_lat_long.json")
	loader := gojsonschema.NewStringLoader(string(body))
	result, err := gojsonschema.Validate(schemaLoader, loader)
	if err != nil {
		panic(err.Error())
	}
	if result.Valid() {
		fmt.Printf("The Document is valid \n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

}
