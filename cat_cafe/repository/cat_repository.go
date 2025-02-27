package repository

import (
	"cat_cafe/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CatRepository struct {
	collection *mongo.Collection
}

func NewCatRepository(db *mongo.Database) *CatRepository {
	return &CatRepository{
		collection: db.Collection("cats"),
	}
}

func (r *CatRepository) CreateCat(cat models.Cat) error {
	_, err := r.collection.InsertOne(context.TODO(), cat)
	return err
}

func (r *CatRepository) GetCats() ([]models.Cat, error) {
	var cats []models.Cat
	fmt.Println("GetCats() ถูกเรียกใช้งาน")

	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("Error ค้นหาข้อมูลจาก MongoDB:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var cat models.Cat
		if err := cursor.Decode(&cat); err != nil {
			fmt.Println("Error Decode JSON:", err)
			return nil, err
		}
		cats = append(cats, cat)
	}

	if len(cats) == 0 {
		fmt.Println("ไม่มีข้อมูลแมวในฐานข้อมูล")
		return []models.Cat{}, nil
	}

	fmt.Println("ดึงข้อมูลแมวสำเร็จ:", cats)
	return cats, nil
}
