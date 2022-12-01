package coda

import "encoding/json"

type CodaResponse struct {
	StatusCode    int           `json:"statusCode"`
	StatusMessage string        `json:"statusMessage"`
	Message       string        `json:"message"`
	Items         []interface{} `json:"items"`
	Href          string        `json:"href"`
	NextPageToken string        `json:"nextPageToken"`
	NextPageLink  string        `json:"nextPageLink"`
}

func GetItemsFromResponse(resp CodaResponse) ([]byte, error) {
	bytes, err := json.MarshalIndent(resp.Items, "", "\t")
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
