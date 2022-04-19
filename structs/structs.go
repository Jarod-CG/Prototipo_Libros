package structs

type CreateAuthorReq struct {
	Name        string `json:"Name,omitempty"`
	DateOfBirth string `json:"DateOfBirth,omitempty"`
}

type CreateAuthorResp struct {
	Success bool   `json:"Success,omitempty"`
	Msg     string `json:"Msg,omitempty"`
}
