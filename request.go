package coda

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func MakeCall(url string, method string, body io.Reader, params map[string]string) ([]byte, error) {
	client := &http.Client{}

	uri := fmt.Sprintf(
		"%s/docs/%s%s",
		os.Getenv("CODA_BASE_URL"),
		os.Getenv("CODA_BASE_DOC"),
		url,
	)

	log.Printf("Sending Query to: %s", uri)

	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CODA_ACCESS_TOKEN")))

	values := req.URL.Query()
	for key, val := range params {
		values.Add(key, val)
	}
	req.URL.RawQuery = values.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}
