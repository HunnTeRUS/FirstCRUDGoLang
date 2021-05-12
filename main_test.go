package main

import (
	"fmt"
	"testing"

	"github.com/oov/dockertest"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
)

func TestMain(m *testing.M) {
	c, err := dockertest.New(dockertest.Config{
		Image: "mongo:latest",
		PortMapping: map[string]string{
			"27017/tcp": "auto",
		},
		Args: []string{"--storageEngine", "wiredTiger"},
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// wait until the container has started listening
	//"127.0.0.1:49157"
	if err = c.Wait(nil); err != nil {
		panic(err)
	}

	session, err := mgo.Dial(c.Mapped["27017/tcp"].String())
	if err != nil {
		panic(err)
	}
	defer session.Close()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:49157"))
	if err != nil {
		panic(err)
	}

	fmt.Println(client)

	type Data struct {
		ID         string
		codeID     string
		channelKey string
	}

	col := session.DB("tokenizationrefusal").C("tokenizationdenial_codes")
	if err = col.Insert(&Data{ID: "AAAAA", codeID: "BBBBB", channelKey: "CCCCCC"}); err != nil {
		panic(err)
	}

	var r Data
	if err = col.Find(nil).One(&r); err != nil {
		panic(err)
	}
}
