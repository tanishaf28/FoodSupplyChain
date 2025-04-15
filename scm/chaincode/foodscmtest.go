package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndQueryProduct(t *testing.T) {
	// Create a mock transaction context
	mockStub := contractapi.NewMockStub("foodscm", new(SmartContract))

	// Init chaincode
	res := mockStub.MockInit("1", [][]byte{})
	assert.Equal(t, int32(contractapi.OK), res.Status, "Init failed")

	// Create Product
	args := [][]byte{
		[]byte("CreateProduct"),
		[]byte("P100"),
		[]byte("Milk"),
		[]byte("Dairy"),
		[]byte("10"),
		[]byte("FarmerX"),
		[]byte("2025-04-14T10:00:00Z"),
	}
	res = mockStub.MockInvoke("1", args)
	assert.Equal(t, int32(contractapi.OK), res.Status, "CreateProduct failed")

	// Query Product
	queryArgs := [][]byte{
		[]byte("QueryProduct"),
		[]byte("P100"),
	}
	res = mockStub.MockInvoke("1", queryArgs)
	assert.Equal(t, int32(contractapi.OK), res.Status, "QueryProduct failed")

	var product Product
	err := json.Unmarshal(res.Payload, &product)
	assert.NoError(t, err)
	assert.Equal(t, "Milk", product.Name)
	assert.Equal(t, "Dairy", product.Category)
	assert.Equal(t, 10, product.Quantity)
	assert.Equal(t, "FarmerX", product.Owner)
	assert.Equal(t, "2025-04-14T10:00:00Z", product.Timestamp)
}

func TestQueryNonExistentProduct(t *testing.T) {
	mockStub := contractapi.NewMockStub("foodscm", new(SmartContract))

	// Init chaincode
	res := mockStub.MockInit("1", [][]byte{})
	assert.Equal(t, int32(contractapi.OK), res.Status)

	// Query a product that doesn't exist
	queryArgs := [][]byte{
		[]byte("QueryProduct"),
		[]byte("P404"),
	}
	res = mockStub.MockInvoke("1", queryArgs)
	assert.NotEqual(t, int32(contractapi.OK), res.Status)
	assert.Contains(t, res.Message, "does not exist")
}
