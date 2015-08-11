package model
import "github.com/bearchinc/datastore-model"

type Post struct {
	db.Model `db:"Post"`
	CarPlate string `json: car_plate`
	Message  string `json: message`
}