// Copyright the Hyperledger Fabric contributors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package extendedsimplecontract

import (
	"errors"
	"fmt"

	"github.com/bestbeforetoday/fabric-contract-api-go/v2/contractapi"
	"github.com/bestbeforetoday/fabric-contract-api-go/v2/internal/functionaltests/contracts/utils"
)

// ExtendedSimpleContract contract for handling writing and reading from the world state
type ExtendedSimpleContract struct {
	contractapi.Contract
}

// Create adds a new key with value to the world state
func (esc *ExtendedSimpleContract) Create(ctx utils.CustomTransactionContextInterface, key string) error {
	existing := ctx.GetCallData()

	if existing != nil {
		return fmt.Errorf("cannot create world state pair with key %s. Already exists", key)
	}

	err := ctx.GetStub().PutState(key, []byte("Initialised"))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

// Update changes the value with key in the world state
func (esc *ExtendedSimpleContract) Update(ctx utils.CustomTransactionContextInterface, key string, value string) error {
	existing := ctx.GetCallData()

	if existing == nil {
		return fmt.Errorf("cannot update world state pair with key %s. Does not exist", key)
	}

	err := ctx.GetStub().PutState(key, []byte(value))

	if err != nil {
		return errors.New("unable to interact with world state")
	}

	return nil
}

// Read returns the value at key in the world state
func (esc *ExtendedSimpleContract) Read(ctx utils.CustomTransactionContextInterface, key string) (string, error) {
	existing := ctx.GetCallData()

	if existing == nil {
		return "", fmt.Errorf("cannot read world state pair with key %s. Does not exist", key)
	}

	return string(existing), nil
}
