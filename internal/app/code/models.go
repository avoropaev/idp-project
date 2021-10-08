package code

import "github.com/google/uuid"

type Code int64

type GuidGenerateResponse struct {
	Token *uuid.UUID `json:"token"`
}

type HashCalcResponse struct {
	Hash *string `json:"hash"`
}
