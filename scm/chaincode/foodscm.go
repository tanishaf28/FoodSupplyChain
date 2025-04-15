package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Product defines the structure for a food product in the supply chain
type Product struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category"`
	Quantity   int    `json:"quantity"`
	Owner      string `json:"owner"`
	Timestamp  string `json:"timestamp"`
}

// SmartContract provides functions for managing the food supply chain
type SmartContract struct {
	contractapi.Contract
}

// InitLedger adds initial data to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	products := []Product{
		{ID: "P001", Name: "Apples", Category: "Fruit", Quantity: 100, Owner: "FarmerA", Timestamp: "2025-04-01T12:00:00Z"},
		{ID: "P002", Name: "Tomatoes", Category: "Vegetable", Quantity: 200, Owner: "FarmerB", Timestamp: "2025-04-01T12:05:00Z"},
	}

	for i, product := range products {
		productAsBytes, err := json.Marshal(product)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState("PRODUCT"+strconv.Itoa(i), productAsBytes)
		if err != nil {
			return fmt.Errorf("failed to put to world state: %v", err)
		}
	}

	return nil
}

// CreateProduct adds a new product to the ledger
func (s *SmartContract) CreateProduct(ctx contractapi.TransactionContextInterface, id, name, category string, quantity int, owner, timestamp string) error {
	product := Product{
		ID:        id,
		Name:      name,
		Category:  category,
		Quantity:  quantity,
		Owner:     owner,
		Timestamp: timestamp,
	}

	productAsBytes, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, productAsBytes)
}

// QueryProduct returns the product stored in the ledger with given id
func (s *SmartContract) QueryProduct(ctx contractapi.TransactionContextInterface, id string) (*Product, error) {
	productAsBytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if productAsBytes == nil {
		return nil, fmt.Errorf("product %s does not exist", id)
	}

	var product Product
	err = json.Unmarshal(productAsBytes, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// QueryAllProducts returns all products in the ledger
func (s *SmartContract) QueryAllProducts(ctx contractapi.TransactionContextInterface) ([]*Product, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var products []*Product
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var product Product
		err = json.Unmarshal(queryResponse.Value, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error create food supply chaincode: %v\n", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %v\n", err)
	}
}
