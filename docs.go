package coda

type CodaDoc struct {
	Id          string               `json:"id"`
	Type        string               `json:"type"`
	Href        string               `json:"href"`
	BrowserLink string               `json:"browserLink"`
	Icon        CodaBasicItem        `json:"icon"`
	Name        string               `json:"name"`
	Owner       string               `json:"owner"`
	OwnerName   string               `json:"ownerName"`
	DocSize     CodaDocSize          `json:"docSize"`
	SourceDoc   CodaBasicItem        `json:"sourceDoc"`
	CreatedAt   string               `json:"createdAt"`
	UpdatedAt   string               `json:"updatedAt"`
	Published   CodaDocPublishedItem `json:"published"`
	Folder      CodaBasicItem        `json:"folder"`
	Workspace   CodaWorkspaceItem    `json:"workspace"`
	WorkspaceId string               `json:"workspaceId"`
	FolderId    string               `json:"folderId"`
}

type CodaDocParams struct {
	IsOwner     bool   `json:"isOwner" url:"isOwner"`
	IsPublished bool   `json:"isPublished"  url:"isPublished"`
	Query       string `json:"query" url:"query"`
	SourceDoc   string `json:"sourceDoc" url:"sourceDoc"`
	IsStarred   bool   `json:"isStarred" url:"isStarred"`
	InGallery   bool   `json:"inGallery" url:"inGallery"`
	WorkspaceId string `json:"workspaceId" url:"workspaceId"`
	FolderId    string `json:"folderId" url:"folderId"`
	Limit       int    `json:"limit" url:"limit"`
	PageToken   string `json:"PageToken" url:"pageToken"`
}

type CodaDocSize struct {
	TotalRowCount     int  `json:"TotalRowCount"`
	TableAndViewCount int  `json:"tableAndViewCount"`
	PageCount         int  `json:"pageCount"`
	OverApiSizeLimit  bool `json:"overApiSizeLimit"`
}

type CodaDocPublishedItem struct {
	Description  string   `json:"description"`
	BrowserLink  string   `json:"browserLink"`
	ImageLink    string   `json:"imageLink"`
	Discoverable bool     `json:"discoverable"`
	EarnCredit   bool     `json:"earnCredit"`
	Mode         string   `json:"mode"`
	Categories   []string `json:"categories"`
}

type CodaWorkspaceItem struct {
	Id             string `json:"id"`
	Type           string `json:"type"`
	OrganizationId string `json:"organizationId"`
	BrowserLink    string `json:"browserLink"`
	Name           string `json:"name"`
}
