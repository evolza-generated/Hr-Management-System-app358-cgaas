package dto

type Salary struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    }

