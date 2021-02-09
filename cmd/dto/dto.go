package dto

import "github.com/Geniuskaa/Task10.1_BGO-3/cmd/card"

type CardDTO struct {
	Error *Error `json:"error"`
	Card card.Card `json:"card"`
}

type Error struct {
	Code int `json:"error_code"`
	Message string `json:"error_msg"`
}
