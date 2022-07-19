package models

type Response struct {
	Code    uint   `json:"code" form:"code" xml:"code"`
	Message string `json:"message" form:"message" xml:"message"`
	Status  string `json:"status" form:"status" xml:"status"`
}

type UsersResponseMany struct {
	Code    uint    `json:"code" form:"code" xml:"code"`
	Message string  `json:"message" form:"message" xml:"message"`
	Status  string  `json:"status" form:"status" xml:"status"`
	Data    []Users `json:"data" form:"data" xml:"data"`
}

type UsersResponseAny struct {
	Code    uint   `json:"code" form:"code" xml:"code"`
	Message string `json:"message" form:"message" xml:"message"`
	Status  string `json:"status" form:"status" xml:"status"`
	Data    Users  `json:"data" form:"data" xml:"data"`
}

type StuffResponseMany struct {
	Code    uint    `json:"code" form:"code" xml:"code"`
	Message string  `json:"message" form:"message" xml:"message"`
	Status  string  `json:"status" form:"status" xml:"status"`
	Data    []Stuff `json:"data" form:"data" xml:"data"`
}

type StuffResponseAny struct {
	Code    uint   `json:"code" form:"code" xml:"code"`
	Message string `json:"message" form:"message" xml:"message"`
	Status  string `json:"status" form:"status" xml:"status"`
	Data    Stuff  `json:"data" form:"data" xml:"data"`
}

type TxResponseAny struct {
	Code    uint   `json:"code" form:"code" xml:"code"`
	Message string `json:"message" form:"message" xml:"message"`
	Status  string `json:"status" form:"status" xml:"status"`
	Data    []Tx   `json:"data" form:"data" xml:"data"`
}

type TxResponseMany struct {
	Code    uint   `json:"code" form:"code" xml:"code"`
	Message string `json:"message" form:"message" xml:"message"`
	Status  string `json:"status" form:"status" xml:"status"`
	Data    []Tx   `json:"data" form:"data" xml:"data"`
}
