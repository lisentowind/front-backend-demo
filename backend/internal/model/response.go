package model

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type TableData struct {
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	List  interface{} `json:"list"`
}

type TableResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data TableData `json:"data"`
}
