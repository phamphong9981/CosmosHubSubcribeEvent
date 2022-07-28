package forwarder

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

var mongoClient, _ = mongo.Connect(ctx, options.Client().ApplyURI(uri))

func saveToMongo(json_data string) {
	var mapData map[string]string
	if err := json.Unmarshal([]byte(json_data), &mapData); err != nil {
		fmt.Println("Unmarshal in MongoDB fail: ", err)
	}
	coll := mongoClient.Database("CosmosHubSubcribeEvent").Collection("undelegate")
	doc := bson.D{{"validator", mapData["validator"]},{"amount",mapData["amount"] },  {"time", mapData["time"]}, {"completion_time", mapData["completion_time"]}, {"delegator", mapData["delegator"]}, {"tx_hash", mapData["tx_hash"]}, {"tx_fee", mapData["tx_fee"]}}
	fmt.Print(doc)
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Println("Write to MongoDB fail: ", err)
	}
}
