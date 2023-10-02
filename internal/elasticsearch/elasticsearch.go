package elasticsearch

type CreateElasticIndexResponse struct {
	Index  string `json:"index"`
	Status string `json:"status"`
}
