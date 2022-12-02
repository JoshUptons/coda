package coda

type CodaTable struct {
	Id            string                 `json:"id"`
	Type          string                 `json:"type"`
	TableType     string                 `json:"tableType"`
	Href          string                 `json:"href"`
	BrowserLink   string                 `json:"browserLink"`
	Name          string                 `jsson:"name"`
	Parent        CodaParentItem         `json:"parent"`
	ParentTable   CodaParentTable        `json:"parentTable"`
	DisplayColumn CodaTableDisplayColumn `json:"displayColumn"`
	RowCount      int                    `json:"rowCount"`
	Sorts         []CodaTableSort        `json:"sorts"`
	Layout        string                 `json:"layout"`
	Filter        CodaTableFilter        `json:"filter"`
	CreatedAt     string                 `json:"createdAt"`
	UpdatedAt     string                 `json:"updatedAt"`
}

type CodaParentTable struct {
	Id          string         `json:"id"`
	Type        string         `json:"type"`
	Href        string         `json:"href"`
	BrowserLink string         `json:"browserLink"`
	Name        string         `json:"name"`
	TableType   string         `json:"tableType"`
	Parent      CodaParentItem `json:"parent"`
}

type CodaTableParams struct {
	Limit      int      `json:"limit"`
	PageToken  string   `json:"pageToken"`
	SortBy     string   `json:"sortBy"`
	TableTypes []string `json:"tableTypes"`
}

type CodaTableDisplayColumn struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Href string `json:"href"`
}

const (
	SortAscending string = "ascending"
	SortDecending        = "descending"
)

type CodaTableSort struct {
	Sorts []struct {
		Column struct {
			Id   string `json:"id"`
			Type string `json:"type"`
			Href string `json:"href"`
		} `json:"column"`
		Direction string `json:"direction"`
	} `json:"sorts"`
}

type CodaTableFilter struct {
	Valid           bool `json:"valid"`
	IsVolatile      bool `json:"IsVolatile"`
	HasUserFormula  bool `json:"hasUserFormula"`
	HasTodayFormula bool `json:"HasTodayFormula"`
	HasNowFormula   bool `json:"hasNowFormula"`
}
