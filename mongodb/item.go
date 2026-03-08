package mongodb

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)


type Item struct {
	ItemName 	string	`json:"item-name",bson:"item-name"`
	Price 		float32	`json:"price",bson:"price"`
	Barcode 	string	`json:"barcode",bson:barcode"`
	Vendor 		string	`json:"vendor",bson:"vendor"`
}

func InsertItem(item Item) error {

	collection := mongoClient.Database(db).Collection(collName)
	inserted, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		panic(err)
	}

	fmt.Println("Inserted record with id: ", inserted.InsertedID)
	return err
}

func GetByName(itemName string) []byte {
	var result Item
		filter := bson.D{{"item-name", itemName}}

	collection := mongoClient.Database(db).Collection(collName)
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}

	return jsonData
}

func GetItems() ([]byte, error) { 
	
	var result []Item
	collection := mongoClient.Database(db).Collection(collName)
	cur, err := collection.Find(context.TODO(),bson.D{})
	if err != nil {
		fmt.Printf("Error retrieving collection:%v",err)
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var item Item
		err := cur.Decode(&item)
		if err != nil {
			fmt.Printf("Error decoding item: %v",err)
			return nil, err
		}
		result = append(result, item)
	}

	cur.Close(context.TODO())

	jsonData, err := json.MarshalIndent(result,"", "    ")
	if err != nil {
		fmt.Printf("Error marshalling item array:%v", err)
		return nil, err
	}

	return jsonData, nil
}

