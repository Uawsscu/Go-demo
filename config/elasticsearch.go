package config

import (
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v7"
)

var ES *elasticsearch.Client

func ConnectElasticSearch() {
	config := elasticsearch.Config{
		Addresses: []string{os.Getenv("ELASTICSEARCH_URL")}, // adress Elasticsearch
	}

	// var Elasticsearch client
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatalf("[elasticsearch] Error creating Elasticsearch client: %v", err)
	}

	// test connect Elasticsearch
	res, err := es.Info()
	if err != nil {
		log.Fatalf("[elasticsearch] Error getting Elasticsearch info: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("[elasticsearch] error: %s", res.Status())
	}

	fmt.Println("[elasticsearch] :", res.Status())
	ES = es
}
