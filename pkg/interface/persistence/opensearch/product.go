package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	opensearch "github.com/opensearch-project/opensearch-go/v2"
	opensearchapi "github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	domainModel "github.com/yomafleet/opensearch-client-test/pkg/domain/model"
	"github.com/yomafleet/opensearch-client-test/pkg/domain/repository"
	"github.com/yomafleet/opensearch-client-test/pkg/interface/persistence/opensearch/model"
)

type OpenSearchProductRepository struct {
	client    *opensearch.Client
	indexName string
}

func (r OpenSearchProductRepository) Create(
	input *domainModel.ProductCreateModel) error {

	if input == nil {

		return nil
	}

	product := model.ConvertFromDomainModel(input)
	productBytes, err := json.Marshal(product)
	if err != nil {

		return err
	}
	document := strings.NewReader(string(productBytes))

	req := opensearchapi.IndexRequest{
		Index:      r.indexName,
		DocumentID: input.ID(),
		Body:       document,
	}
	insertResponse, err := req.Do(context.Background(), r.client)
	fmt.Println(insertResponse)

	return err
}

func (r OpenSearchProductRepository) Search(query string) (
	[]domainModel.ProductGetModel, error) {

	content := strings.NewReader(fmt.Sprintf(`{
	    "query": {
	        "multi_match": {
	            "query": "%s",
	            "fields": ["brand^2", "model", "category"],
				"fuzziness": "AUTO",
				"fuzzy_transpositions": true
	        }
	    },
		"sort": [
			{
				"_score": {
					"order": "desc"
				}
			},
			{
				"quantity": {
					"order": "asc"
				}
			}
		]
	}`, query))
	search := opensearchapi.SearchRequest{

		Index: []string{r.indexName},
		Body:  content,
	}
	searchResponse, err := search.Do(context.Background(), r.client)
	if err != nil {

		return nil, err
	}

	body, err := ioutil.ReadAll(searchResponse.Body)
	if err != nil {

		return nil, err
	}

	bodyMap := make(map[string]interface{})
	if err := json.Unmarshal(body, &bodyMap); err != nil {

		return nil, err
	}

	if v, ok := bodyMap["hits"]; ok {

		if vv, ok := v.(map[string]interface{}); ok {

			if vvv, ok := vv["hits"]; ok {

				if vvvv, ok := vvv.([]interface{}); ok {

					var domainProducts []domainModel.ProductGetModel
					for _, vvvvv := range vvvv {

						jsonBytes, err := json.Marshal(vvvvv)
						if err != nil {

							return nil, err
						}

						product := new(model.ProductGetModel)

						if err := json.Unmarshal(
							jsonBytes, &product); err != nil {

							return nil, err
						}

						if m := product.ConvertToDomainModel(); m != nil {

							domainProducts = append(domainProducts, *m)
						}
					}

					return domainProducts, nil
				}
			}
		}
	}

	return nil, nil
}

func NewOpenSearchProductRepository(
	client *opensearch.Client,
	indexName string) repository.ProductRepositoryInterface {

	return &OpenSearchProductRepository{

		client:    client,
		indexName: indexName,
	}
}
