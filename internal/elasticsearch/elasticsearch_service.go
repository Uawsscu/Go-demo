package elasticsearch

import (
	"context"
	"go-demo/config"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type ElasticsearchService struct {
}

func (cs *ElasticsearchService) CreateElasticIndexCharacter() (*CreateElasticIndexResponse, error) {
	reqBody := `{
		"settings": {
			"number_of_shards": 1,
			"number_of_replicas": 1
		},
		"mappings": {
			"properties": {
				"id": {
					"type": "keyword"
				},
				"name": {
					"type": "text"
				},
				"description": {
					"type": "text"
				},
				"animeID": {
					"type": "keyword"
				},
				"url": {
					"type": "text"
				},
				"createBy": {
					"type": "keyword"
				},
				"updateBy": {
					"type": "keyword"
				},
				"activate": {
					"type": "boolean"
				},
				"createAt": {
					"type": "date"
				},
				"updateAt": {
					"type": "date"
				}
			}
		}
	}`
	elasticsearchRepository := &ElasticsearchRepository{}

	res, err := elasticsearchRepository.CreateElasticIndex(reqBody, "character_index")

	return res, err
}

func (es *ElasticsearchService) InsertDocument(index, documentID string, documentBody string) error {
	// Create the request for indexing a document
	reqBodyReader := strings.NewReader(documentBody)

	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       reqBodyReader,
		Refresh:    "true",
	}

	// Perform the request
	res, err := req.Do(context.Background(), config.ES)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
