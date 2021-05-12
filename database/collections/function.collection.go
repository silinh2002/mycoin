package collection

import (
	"context"
	"log"
	"mycoin/database/data_context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	*mongo.Collection
}

func GetCollection(collectionName string) *mongo.Collection {
	return data_context.GetInstance().DB.Collection(collectionName)
}

func (collection Collection) GetList(operations []bson.M) []bson.Raw {
	cursor, err := collection.Aggregate(context.TODO(), operations)
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.Raw
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func (collection Collection) GetLastRecord() []bson.Raw {

	sort := bson.M{"$sort": bson.M{"_id": -1}}
	limit := bson.M{"$limit": 1}
	conditions := []bson.M{sort, limit}

	cursor, err := collection.Aggregate(context.TODO(), conditions)
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.Raw
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func (collection Collection) FindByLambda(condition interface{}) []bson.Raw {
	//opts := options.Find().SetSort(bson.D{{"age", 1}})
	cursor, err := collection.Find(context.TODO(), condition)
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.Raw
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func (collection Collection) CreateByLambda(lambda interface{}) (*mongo.InsertOneResult, error) {
	//fmt.Println("lambda:", lambda)
	cursor, err := collection.InsertOne(context.TODO(), lambda)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

func (collection Collection) UpdateByLambda(condition interface{}, lambda interface{}) (*mongo.UpdateResult, error) {
	cursor, err := collection.UpdateOne(context.TODO(), condition, bson.D{{"$set", lambda}})
	if err != nil {
		return nil, err

	}
	return cursor, nil
}

// func (collection Collection) UpdateMany(condition interface{}, lambda interface{}) (*mongo.UpdateResult, error) {
// 	cursor, err := collection.UpdateMany(context.TODO(), condition, bson.D{{"$set", lambda}})
// 	if err != nil {
// 		return nil, err

// 	}
// 	return cursor, nil
// }
