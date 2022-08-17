package services

import (
	"context"

	"github.com/Gprisco/decanto-winefamily-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetWinefamilies(page int64, limit int64) []models.Winefamily {
	findOptions := options.Find()

	findOptions.SetLimit(limit)
	findOptions.SetSkip((page - 1) * limit)

	var winefamilies []models.Winefamily
	cursor, err := models.WinefamilyCollection.Find(context.TODO(), bson.D{}, findOptions)

	if err != nil {
		panic(err)
	}

	cursor.All(context.TODO(), &winefamilies)

	return winefamilies
}

func GetWinefamily(id primitive.ObjectID) models.Winefamily {
	var winefamily models.Winefamily
	err := models.WinefamilyCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&winefamily)

	if err != nil {
		panic(err)
	}

	return winefamily
}
