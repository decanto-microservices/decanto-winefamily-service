package models

import (
	"github.com/Gprisco/decanto-winefamily-service/db"
	"github.com/Gprisco/decanto-winefamily-service/env"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Winefamily struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Winefamily string             `bson:"winefamily" json:"winefamily"`
}

var WinefamilyCollection *mongo.Collection = db.GetInstance().Client().Database(env.GetInstance().DB).Collection("a_winefamilies")
