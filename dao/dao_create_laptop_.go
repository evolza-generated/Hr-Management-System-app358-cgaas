package dao

import (
    "context"
	"Hr-Management-System/dbConfig"
	"Hr-Management-System/dto"

)

func DB_CreateLaptop (application *dto.Laptop) error {

	_, err := dbConfig.DATABASE.Collection("Laptops").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}