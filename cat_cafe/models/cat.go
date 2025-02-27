package models

type Cat struct {
	Name         string `json:"name" bson:"name"`
	Breed        string `json:"breed" bson:"breed"`
	Age          int    `json:"age" bson:"age"`
	FavoriteFood string `json:"favorite_food" bson:"favorite_food"`
	Status       string `json:"status" bson:"status"`
}
