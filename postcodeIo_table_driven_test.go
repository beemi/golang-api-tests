package main

import (
	"encoding/json"
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
//Validate Json schema validation from file from local disk using gojsonschema package
//Assert response using testify package
func TestTablePostCodeLatLong(t *testing.T) {

	testCases := []struct {
		postCode  string
		want      string
		longitude float64
		latitude  float64
	}{
		{"RM17 6EY", "RM17 6EY", 0.342289, 51.476936},
		{"IG1 2FJ", "IG1 2FJ", 0.067261, 51.557962},
	}

	for _, tc := range testCases {
		url := config.PostCodeIOEndPoint() + "/postcodes/" + tc.postCode
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

		var postCodeResponse PostCodeIO
		//Json Unmarshal
		err = json.Unmarshal([]byte(string(body)), &postCodeResponse)
		assert.NoError(t, err)
		//validate response objects with assertions
		assert.Equal(t, 200, postCodeResponse.Status)
		assert.Equal(t, tc.want, postCodeResponse.Result.Postcode, "result.postcode value is wrong")
		assert.Equal(t, tc.latitude, postCodeResponse.Result.Latitude, "result.latitude value is wrong")
		assert.Equal(t, tc.longitude, postCodeResponse.Result.Longitude, "result.longitude value is wrong")
	}
}
