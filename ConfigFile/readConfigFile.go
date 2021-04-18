package configFile

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mongo       string
	Collections []string
}

func ReadConfigFile() Config {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var config Config
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	result := Config{
		Mongo:       config.Mongo,
		Collections: config.Collections,
	}

	fmt.Println(result)

	return result
}
