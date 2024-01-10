package util

type ResponseData struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

type BasicErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}
