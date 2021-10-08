package models

type HashCalcRequest struct {
	Code int64 `json:"code"`
}

type HashCalcResponse struct {
	Hash string `json:"hash"`
}
