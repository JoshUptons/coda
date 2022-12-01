package coda

type CodaItem struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	TableType   string `json:"tableType"`
	Href        string `json:"href"`
	BrowserLink string `json:"browserLink"`
	Name        string `json:"name"`
	Parent      CodaBasicItem
}

type CodaBasicItem struct {
	Id          string         `json:"id"`
	Type        string         `json:"type"`
	Href        string         `json:"href"`
	BrowserLink string         `json:"browserLink"`
	Name        string         `json:"name"`
	Parent      *CodaBasicItem `json:"parent"`
}
