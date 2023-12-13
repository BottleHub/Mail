package models

type Email struct {
	Address string `json:"address" validate:"required" bson:"address"`
}
