package main

import (
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

//TestPostCodeLatLong test validate lat long validation
//Validate Json schema validation using gojsonschema package
//Assert response using testify package
func TestPostCodeLatLong(t *testing.T) {
	url := config.PostCodeIOEndPoint() + "/postcodes/RM17%206EY"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	//assert ensure no error
	assert.NoError(t, err)

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	assert.NoError(t, err)

	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode, "Get Postcode lat long Api failed")
	body, err := ioutil.ReadAll(res.Body)
	log.Printf("Response Body: \n %s", string(body))

	// JSON schema validation
	dir, err := os.Getwd()
	assert.NoError(t, err)
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + dir + "/schemas/postcode_lat_long.json")
	loader := gojsonschema.NewStringLoader(string(body))
	result, err := gojsonschema.Validate(schemaLoader, loader)
	assert.NoError(t, err)

	if result.Valid() {
		log.Printf("The Document is valid \n")
	} else {
		log.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			log.Printf("- %s\n", desc)
		}
		t.Fail()
	}
}
