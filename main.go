package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"

	opensearch "github.com/opensearch-project/opensearch-go/v2"
	opensearchapi "github.com/opensearch-project/opensearch-go/v2/opensearchapi"
)

func main() {

	// Initialize the client with SSL/TLS enabled.
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // For testing only. Use certificate for validation.
		},
		Addresses: []string{"https://localhost:9200"},
		Username:  "admin", // For testing only. Don't store credentials in code.
		Password:  "admin",
	})
	if err != nil {

		panic(err)
	}

	// Print OpenSearch version information on console.
	fmt.Println(client.Info())

	// Define index mapping.
	mapping := strings.NewReader(`{
		"settings": {
			"number_of_shards": 1,
			"number_of_replicas": 1
		},
		"mappings": {
			"properties": {
				"brand": {
					"type": "text"
				},
				"model": {
					"type": "text"
				},
				"category": {
					"type": "text"
				}
			}
		}
	}`)

	// Create an index with non-default settings.
	indexName := "products-brand-model-category-index"
	createIndex := opensearchapi.IndicesCreateRequest{

		Index: indexName,
		Body:  mapping,
	}
	ctx := context.Background()
	createIndexResponse, err := createIndex.Do(ctx, client)
	if err != nil {

		panic(err)
	}
	fmt.Println(createIndexResponse)

	// Add a document to the index.
	// document := strings.NewReader(`{
	//     "name": "iPhone",
	//     "model": "13",
	//     "quantity": 16
	// }`)

	// req := opensearchapi.IndexRequest{
	// 	Index: indexName,
	// 	Body:  document,
	// }
	// insertResponse, err := req.Do(ctx, client)
	// if err != nil {

	// 	panic(err)
	// }
	// fmt.Println(insertResponse)

	// Search for the document.
	// content := strings.NewReader(`{
	//     "query": {
	//         "multi_match": {
	//             "query": "phone",
	//             "fields": ["name^2", "model"],
	// 			"fuzziness": "AUTO",
	// 			"fuzzy_transpositions": true
	//         }
	//     },
	// 	"sort": [
	// 		{
	// 			"_score": {
	// 				"order": "desc"
	// 			}
	// 		},
	// 		{
	// 			"quantity": {
	// 				"order": "asc"
	// 			}
	// 		}
	// 	]
	// }`)
	// search := opensearchapi.SearchRequest{

	// 	Index: []string{indexName},
	// 	Body:  content,
	// }

	// searchResponse, err := search.Do(ctx, client)
	// if err != nil {

	// 	panic(err)
	// }
	// fmt.Println(searchResponse)
}
