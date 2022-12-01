package coda

import (
	"encoding/json"
	"fmt"
)

type CodaTable struct {
	Id            string            `json:"id"`
	Type          string            `json:"type"`
	TableType     string            `json:"tableType"`
	Href          string            `json:"href"`
	BrowserLink   string            `json:"browserLink"`
	Name          string            `jsson:"name"`
	Parent        CodaBasicItem     `json:"parent"`
	ParentTable   *CodaTable        `json:"parentTable"`
	DisplayColumn CodaDisplayColumn `json:"displayColumn"`
	RowCount      int               `json:"rowCount"`
	Sorts         []CodaSort        `json:"sorts"`
	Layout        string            `json:"layout"`
	Filter        CodaFilter        `json:"filter"`
	CreatedAt     string            `json:"createdAt"`
	UpdatedAt     string            `json:"updatedAt"`
}

func GetTable(table string) (CodaTable, error) {
	var response CodaTable

	resp, err := MakeCall(fmt.Sprintf("/tables/%s", table), "GET", nil, map[string]string{
		"useColumnNames": "true",
	})
	if err != nil {
		return CodaTable{}, err
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		return CodaTable{}, err
	}

	return response, nil
}

func GetTables() ([]CodaTable, error) {
	var response []CodaTable

	return response, nil
}
