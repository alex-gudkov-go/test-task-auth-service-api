package models

type RefreshToken struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	Value  string `json:"value"`
}
