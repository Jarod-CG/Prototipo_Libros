package structs

type response struct {
	Success bool   `json:"Success,omitempty"`
	Msg     string `json:"Msg,omitempty"`
}

//create author
type CreateAuthorReq struct {
	Name        string `json:"Name,omitempty"`
	DateOfBirth string `json:"DateOfBirth,omitempty"`
}

type CreateAuthorResp struct {
	response
}

//create book
type CreateBookReq struct {
	AuthorName string  `json:"Name,omitempty"`
	BookName   string  `json:"DateOfBirth,omitempty"`
	Topic      string  `json:"Topic,omitempty"`
	Date       string  `json:"Date,omitempty"`
	Price      float64 `json:"Price,omitempty"`
	Quantity   int     `json:"Quantity,omitempty"`
}

type CreateBookResp struct {
	response
}

//create reader
type CreateReaderReq struct {
	ReaderName string `json:"ReaderName,omitempty"`
}

type CreateReaderResp struct {
	response
}

//order book
type OrderBookReq struct {
	ReaderName string `json:"ReaderName,omitempty"`
	BookName   string `json:"BookName,omitempty"`
}

type OrderBookResp struct {
	response
}

//buy book
type BuyBookReq struct {
	ReaderName string `json:"ReaderName,omitempty"`
	BookName   string `json:"BookName,omitempty"`
}

type BuyBookResp struct {
	response
}

//watch author
type WatchAuthorReq struct {
	ReaderName string `json:"ReaderName,omitempty"`
	AuthorName string `json:"AuthorName,omitempty"`
}

type WatchAuthorResp struct {
	response
}

//watch book
type WatchBookReq struct {
	ReaderName string `json:"ReaderName,omitempty"`
	BookName   string `json:"BookName,omitempty"`
}

type WatchBookResp struct {
	response
}

type CleanDBReq struct {
}

type CleanDBResp struct {
	response
}
