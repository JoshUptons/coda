package coda

type CodaColumn struct {
	Id           string           `json:"id"`
	Type         string           `json:"type"`
	Href         string           `json:"href"`
	Name         string           `json:"name"`
	Display      bool             `json:"display"`
	Calculated   bool             `json:"calculated"`
	Formula      string           `json:"formula"`
	DefaultValue string           `json:"defaultValue"`
	Format       CodaColumnFormat `json:"format"`
}

type CodaColumnFormat struct {
	Type      string `json:"type"`
	IsArray   bool   `json:"isArray"`
	Label     string `json:"label"`
	DisableIf string `json:"disableIf"`
	Action    string `json:"action"`
}
