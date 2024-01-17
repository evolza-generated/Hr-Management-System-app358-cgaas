package dao

import (
	"Hr-Management-System/dbConfig"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_DeleteEmployees (objectid string)  error {

  
        result, err := dbConfig.DATABASE.Collection("Employeess").DeleteOne(context.Background(), bson.M{"objectid": objectid})
        if err != nil {
            return err
        }
        if result.DeletedCount < 1 {
            return errors.New("Specified Id not found!")
        }
        return nil
    
}

