package code

import "github.com/google/uuid"

type Code int

type GuidGenerateResponse struct {
	Token *uuid.UUID `json:"token"`
}

type HashCalcResponse struct {
	Hash *string `json:"hash"`
}
