package elasticsearch

import (
	"context"
	"fmt"
	"go-demo/config"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type ElasticsearchRepository struct{}

func (es *ElasticsearchRepository) CreateElasticIndex(reqBody string, index string) (*CreateElasticIndexResponse, error) {
	// Create the request for creating the index
	reqBodyReader := strings.NewReader(reqBody)

	// Create the request for creating the index
	req := esapi.IndicesCreateRequest{
		Index: index,
		Body:  reqBodyReader, // Use the io.Reader here
	}

	// Perform the request
	res, err := req.Do(context.Background(), config.ES)

	if err != nil {
		log.Fatalf("Error creating index: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Elasticsearch error: %s", res.Status())
	}

	fmt.Printf("Index '%s' created successfully\n", index)

	return &CreateElasticIndexResponse{
		Index:  index,
		Status: "success",
	}, err

}
