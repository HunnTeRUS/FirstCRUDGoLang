package test_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMain2(m *testing.T) {
	// var db *mongo.Client

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not open a pool: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://127.0.0.1:%s", resource.GetPort("27017/tcp"))))
	if err != nil {
		log.Fatalf("Could not create a client: %s", err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Could not ping on database: %s", err)
	}

	type Data struct {
		ID         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
		CodeID     string             `bson:"codeId" json:"codeId"`
		ChannelKey string             `bson:"channelKey" json:"channelKey"`
	}

	collection := client.Database("dbOdontooth").Collection("dentists")

	_, err = collection.InsertOne(context.Background(), &Data{CodeID: "BBBBB", ChannelKey: "CCCCCC"})
	if err != nil {
		log.Fatal(err)
	}

	var results []*Data

	cur, err := collection.Find(context.TODO(), primitive.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var elem Data
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	err = resource.Close()
	if err != nil {
		log.Fatalf("Could not stop docker container: %s", err)
	}
}

func TestSomething(t *testing.T) {
	// db.Query()
}
