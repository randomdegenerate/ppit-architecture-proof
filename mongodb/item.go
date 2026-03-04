package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
)


type Item struct {
	ItemName 	string	`json:"item-name",bson:"item-name"`
	price 		float32	`json:"price",bson:"price"`
	barcode 	string	`json:"barcode",bson:"barcode"`
	vendor 		string	`json:"vendor",bson:"vendor"`
}

func InsertItem(item Item) error {

	collection := mongoClient.Database(db).Collection(collName)
	inserted, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted record with id: ", inserted.InsertedID)
	return err
}

func FindByName(itemName string) []byte {
	var result Item

		filter := bson.D{{"item-name", itemName}}

	collection := mongoClient.Database(db).Collection(collName)
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return jsonData
}

