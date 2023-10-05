package elasticsearch

import (
	"context"
	"errors"
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
		log.Printf("Error creating index: %v", err)
		return nil, errors.New(err.Error())
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Elasticsearch error: %s", res.Status())
		return nil, errors.New("index is exist")
	}

	fmt.Printf("Index '%s' created successfully\n", index)

	return &CreateElasticIndexResponse{
		Index: index,
	}, err

}
