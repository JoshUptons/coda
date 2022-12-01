package coda

type CodaDisplayColumn struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Href string `json:"href"`
}

type CodaSort struct {
	Column struct {
		Id   string `json:"id"`
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"column"`
	Direction string `json:"direction"`
}

type CodaFilter struct {
	Valid           bool `json:"valid"`
	IsVolatile      bool `json:"isVolatile"`
	HasUserFormula  bool `json:"hasUserFormula"`
	HasTodayFormula bool `json:"hasTodayFormula"`
	HasNowFormula   bool `json:"hasNowFormula"`
}
