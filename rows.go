package coda

import (
	"encoding/json"
	"fmt"
)

type CodaRow struct {
	Id          string                 `json:"id"`
	Type        string                 `json:"type"`
	Href        string                 `json:"href"`
	Name        string                 `json:"name"`
	Index       int                    `json:"index"`
	BrowserLink string                 `json:"browserLink"`
	CreatedAt   string                 `json:"createdAt"`
	UpdatedAt   string                 `json:"updatedAt"`
	Values      map[string]interface{} `json:"values"`
	Parent      CodaBasicItem          `json:"parent"`
}

type CodaRowValues struct {
}

func GetRows(table string) ([]CodaRow, error) {
	var response []CodaRow

	resp, err := MakeCall(fmt.Sprintf("/tables/%s/rows", table), "GET", nil, map[string]string{
		"useColumnNames": "true",
	})
	if err != nil {
		return response, err
	}

	var codaResponse CodaResponse
	err = json.Unmarshal(resp, &codaResponse)
	if err != nil {
		return response, err
	}

	rows, err := GetItemsFromResponse(codaResponse)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rows, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
