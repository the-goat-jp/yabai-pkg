package model

type JWTData struct {
	User
	RefID string `json:"ref"`
}
