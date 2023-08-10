package model

type MsgStruct struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type JsonStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}
