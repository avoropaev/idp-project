package models

type GuidGenerateRequest struct {
	Code int64 `json:"code"`
}

type GuidGenerateResponse struct {
	Token string `json:"token"`
}
