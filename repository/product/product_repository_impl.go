package product

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/thanhvdt/vcs-week2/model"
	"gorm.io/gorm"
	"io"
	"log"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
	Es *elasticsearch.Client
}

func NewProductRepository(db *gorm.DB, es *elasticsearch.Client) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{Db: db, Es: es}
}

func (p *ProductRepositoryImpl) GetProductsWithSupplier() ([]model.Product, error) {
	var products []model.Product
	if err := p.Db.Preload("Supplier").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepositoryImpl) GetProductAboveAveragePrice() ([]model.Product, error) {
	var products []model.Product
	if err := p.Db.Where("unit_price > (?)", p.Db.Table("products").Select("AVG(unit_price)")).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepositoryImpl) UpdateInElasticSearch(docID string, updateFields map[string]interface{}) error {
	updateReq := map[string]interface{}{
		"doc": updateFields,
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(updateReq); err != nil {
		log.Printf("Error encoding update request: %s", err)
		return err
	}

	res, err := p.Es.Update(
		"search-product",
		docID,
		bytes.NewReader(buf.Bytes()),
		p.Es.Update.WithContext(context.Background()),
	)
	if err != nil {
		log.Printf("Error performing update request: %s", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %s", err)
		}
	}(res.Body)

	if res.IsError() {
		log.Printf("Error response from Elasticsearch: %s", res.String())
		return fmt.Errorf("error updating document ID=%s, response: %s", docID, res.String())
	}

	log.Printf("Successfully updated document ID=%s, response status: %s", docID, res.Status())

	return nil
}

func (p *ProductRepositoryImpl) CreateInElasticSearch(document *map[string]interface{}) (string, error) {
	var buf bytes.Buffer
	fmt.Println("document in repo", document)
	if err := json.NewEncoder(&buf).Encode(document); err != nil {
		log.Printf("Error encoding document: %s", err)
		return "", err
	}

	// Perform the create operation without specifying the document ID
	res, err := p.Es.Index(
		"search-product",
		&buf,
	)
	if err != nil {
		log.Printf("Error performing create operation: %s", err)
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %s", err)
		}
	}(res.Body)

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s", err)
			return "", err
		}
		errMsg := fmt.Sprintf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		log.Printf(errMsg)
		return "", fmt.Errorf(errMsg)
	}

	// Extract the generated document ID from the response
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body to get the document ID: %s", err)
		return "", err
	}
	docID, ok := r["_id"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get the document ID from the Elasticsearch response")
	}

	log.Printf("Document created successfully in index 'search-product' with ID=%s\n", docID)
	return docID, nil
}

func (p *ProductRepositoryImpl) GetAllInElasticSearch() (map[string]interface{}, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
		return nil, err
	}
	res, err := p.Es.Search(
		p.Es.Search.WithContext(context.Background()),
		p.Es.Search.WithIndex("search-product"),
		p.Es.Search.WithBody(&buf),
		p.Es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error closing response body: %s", err)
		}
	}(res.Body)

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
			return nil, err
		} else {
			// Print the error message from Elasticsearch
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	// Decode the response body to a map or a specific struct
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return nil, err
	}

	_, err = json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %s", err)
		return nil, err
	}

	return r, nil
}

func (p *ProductRepositoryImpl) GetByIdInElasticSearch(docID string) (map[string]interface{}, error) {
	res, err := p.Es.Get("search-product", docID)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error closing response body: %s", err)
		}
	}(res.Body)

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
			return nil, err
		} else {
			// Print the error message from Elasticsearch
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	// Decode the response body to a map or a specific struct
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return nil, err
	}

	_, err = json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %s", err)
		return nil, err
	}

	return r, nil
}

func (p *ProductRepositoryImpl) DeleteInElasticSearch(docID string) error {
	res, err := p.Es.Delete("search-product", docID)
	if err != nil {
		log.Printf("Error deleting document ID=%s: %s", docID, err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %s", err)
		}
	}(res.Body)
	if res.IsError() {
		log.Printf("Error response from Elasticsearch: %s", res.String())
		return fmt.Errorf("error deleting document ID=%s, response: %s", docID, res.String())
	}
	return nil
}
