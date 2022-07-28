package backend

import (
	"context"
	"encoding/json"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

var mongoClient, _ = mongo.Connect(ctx, options.Client().ApplyURI(uri))

func findAll(view_more string) []string {
	var results []string
	skip, err := strconv.ParseInt(view_more, 10, 0)
	if err != nil {
		panic(err)
	}
	coll := mongoClient.Database("CosmosHubSubcribeEvent").Collection("undelegate")
	options := options.Find()
	options.SetLimit(3)
	options.SetSkip(skip * 3)
	options.SetSort(bson.D{{"time", -1}})

	cursor, err := coll.Find(context.TODO(), bson.D{}, options)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		panic(err)
	}
	var bson_data []bson.M
	if err = cursor.All(context.TODO(), &bson_data); err != nil {
		panic(err)
	}
	for _, element := range bson_data {
		output, err := json.Marshal(element)
		if err != nil {
			panic(err)
		}
		results = append(results, string(output))
	}
	return results
}

func findByDelegator(delegator string) []string {
	var results []string
	coll := mongoClient.Database("CosmosHubSubcribeEvent").Collection("undelegate")
	options := options.Find()
	options.SetLimit(1)
	cursor, err := coll.Find(context.TODO(), bson.D{{"delegator", delegator}}, options)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		panic(err)
	}
	var bson_data []bson.M
	if err = cursor.All(context.TODO(), &bson_data); err != nil {
		panic(err)
	}
	for _, element := range bson_data {
		output, err := json.Marshal(element)
		if err != nil {
			panic(err)
		}
		results = append(results, string(output))
	}
	return results
}

func getUnbondFromDelegator(delegator string, view_more string) []string {

	if delegator == "" {
		return nil
	}
	if delegator == "all" {

		return findAll(view_more)
	}
	return findByDelegator(delegator)
}
