package models

type GUIDGenerateRequest struct {
	Code int64 `json:"code"`
}

type GUIDGenerateResponse struct {
	Token string `json:"token"`
}
