package models

type ResponseJson struct {
	Status int         `json:"_"`
	Code   int         `json:"code,omitempty"`
	Msg    string      `json:"string,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
