package coda

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

const BaseUrl = "https://coda.io/apis/v1/docs"

type Client struct {
	HttpClient *http.Client
	Token      string
}

func Default(token string) *Client {
	return &Client{
		HttpClient: &http.Client{},
		Token:      token,
	}
}

func (c *Client) MakeCall(url, method string, body io.Reader, params interface{}) ([]byte, error) {
	var response []byte
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return response, nil
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	values := req.URL.Query()
	err = decoder.Decode(&params, values)
	if err != nil {
		log.Fatalf("error decoding params into query, %v", err)
	} else {
		log.Println("set parameters: ", params)
	}

	req.URL.RawQuery = values.Encode()

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) GetTables(doc, table string, params CodaTableParams) ([]CodaTable, error) {
	var response []CodaTable

	url := fmt.Sprintf("%s/%s/%s", BaseUrl, doc, table)

	c.MakeCall(url, "GET", nil, params)

	return response, nil
}

func (c *Client) GetDoc(doc string, params CodaDocParams) (CodaDoc, error) {
	var response CodaDoc

	url := fmt.Sprintf("%s/%s", doc)

	c.MakeCall(url, "GET", nil, params)

	return response, nil
}

func (c *Client) ListDocs(params CodaDocParams) ([]CodaDoc, error) {
	var response []CodaDoc

	c.MakeCall(BaseUrl, "GET", nil, params)

	return response, nil
}
