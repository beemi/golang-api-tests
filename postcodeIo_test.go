package main

import (
	"encoding/json"
	"fmt"
	"github.com/beemi/postcode-io-tests-golang/internal/config"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

// PostCodeIO payload structure for response
type PostCodeIO struct {
	Status int    `json:"status"`
	Result Result `json:"result"`
}

// Result is the Postcode Io
type Result struct {
	Postcode                  string      `json:"postcode"`
	Quality                   int64       `json:"quality"`
	Eastings                  int64       `json:"eastings"`
	Northings                 int64       `json:"northings"`
	Country                   string      `json:"country"`
	NhsHa                     string      `json:"nhs_ha"`
	Longitude                 float64     `json:"longitude"`
	Latitude                  float64     `json:"latitude"`
	EuropeanElectoralRegion   string      `json:"european_electoral_region"`
	PrimaryCareTrust          string      `json:"primary_care_trust"`
	Region                    string      `json:"region"`
	Lsoa                      string      `json:"lsoa"`
	Msoa                      string      `json:"msoa"`
	Incode                    string      `json:"incode"`
	Outcode                   string      `json:"outcode"`
	ParliamentaryConstituency string      `json:"parliamentary_constituency"`
	AdminDistrict             string      `json:"admin_district"`
	Parish                    string      `json:"parish"`
	AdminCounty               interface{} `json:"admin_county"`
	AdminWard                 string      `json:"admin_ward"`
	Ced                       interface{} `json:"ced"`
	Ccg                       string      `json:"ccg"`
	Nuts                      string      `json:"nuts"`
	Codes                     Codes       `json:"codes"`
}

// Codes postcode response codes
type Codes struct {
	AdminDistrict             string `json:"admin_district"`
	AdminCounty               string `json:"admin_county"`
	AdminWard                 string `json:"admin_ward"`
	Parish                    string `json:"parish"`
	ParliamentaryConstituency string `json:"parliamentary_constituency"`
	Ccg                       string `json:"ccg"`
	Ced                       string `json:"ced"`
	Nuts                      string `json:"nuts"`
}

// TestPostCodeLatLong test validate lat long validation
// Validate Json schema validation from file from local disk using gojsonschema package
// Assert response using testify package
func TestPostCodeLatLong(t *testing.T) {
	testPostcode := "RM17 6EY"
	url := config.PostCodeIOEndPoint() + "/postcodes/%s"
	method := "GET"

	// add postcode to url
	url = fmt.Sprintf(url, testPostcode)

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	//assert ensure no error
	assert.NoError(t, err)

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	assert.NoError(t, err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	assert.Equal(t, 200, res.StatusCode, "Get Postcode lat long Api failed")
	body, err := ioutil.ReadAll(res.Body)
	log.Printf("Response Body: \n %s", string(body))

	//assert ensure no error
	assert.NoError(t, err)
	var postCodeResponse PostCodeIO
	//Json Unmarshal
	err = json.Unmarshal([]byte(body), &postCodeResponse)
	assert.NoError(t, err)
	//validate response objects with assertions
	assert.Equal(t, 200, postCodeResponse.Status)
	assert.Equal(t, testPostcode, postCodeResponse.Result.Postcode, "result.postcode value is wrong")
}
