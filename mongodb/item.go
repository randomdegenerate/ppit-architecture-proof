package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	//"encoding/json"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Item struct {
	ID bson.ObjectID `json: "_id,omitempty" bson:"_id,omitempty"`
	item_name string `json:"item_name"`
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

func InsertMany(items []Item) error {
		
	collection := mongoClient.Database(db).Collection(collName)
	result, err := collection.InsertMany(context.TODO(), items)
	if err != nil {
		log.Fatal(err)	
	}

	fmt.Println(result)
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

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}

func FindById(id string) []byte {
	var result Item

	filter := bson.D{{"_id", id}}

	collection := mongoClient.Database(db).Collection(collName)
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}
