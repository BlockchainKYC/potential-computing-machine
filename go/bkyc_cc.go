/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// KYCreg provides functions for managing a User reg
type KYCreg struct {
	contractapi.Contract
}

// User describes basic details of what makes up a User
type User struct {
	FirstName   	string `json:"firstName"`
	LastName    	string `json:"lastName"`
	Gender  		string `json:"gender"`
	Email     		string `json:"email"`
	PhoneNumber		string `json:"phoneNumber"`
	Address 		string `json:"address"`
	Key			 	string `json:"registrationId"`
	DocHash 		string `json:"docHash"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *User
}

// InitLedger adds a base set of users to the ledger
func (s *KYCreg) InitLedger(ctx contractapi.TransactionContextInterface) error {
	users := []User{
		// User{Name: "Kavitha", Gender: "Female", Pan: "QWERT1234Y", Aadhar: "111122223333", Address: "Nit Warangal,506009"},
		// User{Name: "Shyam", Gender: "Male", Pan: "ASDFG5678H", Aadhar: "222233334444", Address: "Nit Surathkal,511283"},
		// User{Name: "Anand", Gender: "Male", Pan: "ZXCVB2345N", Aadhar: "333344445555", Address: "Nit Trichy,543213"},
		// User{Name: "Nithya", Gender: "Female", Pan: "MNBVC9876X", Aadhar: "444455556666", Address: "Nit Bhopal,583214"},
	}

	for i, user := range users {
		userAsBytes, _ := json.Marshal(user)
		err := ctx.GetStub().PutState("USER"+strconv.Itoa(i), userAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

// CreateUser adds a new user to the world state with given details
func (s *KYCreg) CreateUser(ctx contractapi.TransactionContextInterface, userNumber string, firstName string, lastName string, gender string, email string, phoneNumber string, address string) error {
	user := User{
		FirstName: 			firstName,
		LastName: 			lastName,
		Gender:  			gender,
		Address: 			address,
		PhoneNumber:     	phoneNumber,
		Email:  			email,
		
	}


	fmt.Println("userNumber", user);
	userAsBytes, _ := json.Marshal(user)

	return ctx.GetStub().PutState(userNumber, userAsBytes)
}

// QueryUser returns the user stored in the world state with given id
func (s *KYCreg) QueryUser(ctx contractapi.TransactionContextInterface, userNumber string) (*User, error) {
	userAsBytes, err := ctx.GetStub().GetState(userNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if userAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", userNumber)
	}

	user := new(User)
	_ = json.Unmarshal(userAsBytes, user)

	return user, nil
}

// QueryAllUsers returns all users found in world state
func (s *KYCreg) QueryAllUsers(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		user := new(User)
		_ = json.Unmarshal(queryResponse.Value, user)

		queryResult := QueryResult{Key: queryResponse.Key, Record: user}
		results = append(results, queryResult)
	}

	return results, nil
}

// ChangeUserAddress updates the address field of user with given id in world state
func (s *KYCreg) ChangeUserAddress(ctx contractapi.TransactionContextInterface, userNumber string, newAddress string) error {
	user, err := s.QueryUser(ctx, userNumber)

	if err != nil {
		return err
	}

	user.Address = newAddress

	userAsBytes, _ := json.Marshal(user)

	return ctx.GetStub().PutState(userNumber, userAsBytes)
}

func (s *KYCreg) ChangeUserFirstName(ctx contractapi.TransactionContextInterface, userNumber string, newFirstName string) error {
	user, err := s.QueryUser(ctx, userNumber)

	if err != nil {
		return err
	}

	user.FirstName = newFirstName

	userAsBytes, _ := json.Marshal(user)

	return ctx.GetStub().PutState(userNumber, userAsBytes)
}

func (s *KYCreg) ChangeUserLastName(ctx contractapi.TransactionContextInterface, userNumber string, newLastName string) error {
	user, err := s.QueryUser(ctx, userNumber)

	if err != nil {
		return err
	}

	user.LastName = newLastName

	userAsBytes, _ := json.Marshal(user)

	return ctx.GetStub().PutState(userNumber, userAsBytes)
}

func (s *KYCreg) ChangeGender(ctx contractapi.TransactionContextInterface, userNumber string, newGender string) error {
	user, err := s.QueryUser(ctx, userNumber)

	if err != nil {
		return err
	}

	user.Gender = newGender

	userAsBytes, _ := json.Marshal(user)
	// print newgender

	fmt.Println(newGender)
	return ctx.GetStub().PutState(userNumber, userAsBytes)
}

func (s *KYCreg) ChangeUserPhoneNumber(ctx contractapi.TransactionContextInterface, userNumber string, newPhoneNumber string) error {
	user, err := s.QueryUser(ctx, userNumber)

	if err != nil {
		return err
	}

	user.PhoneNumber = newPhoneNumber

	userAsBytes, _ := json.Marshal(user)

	return ctx.GetStub().PutState(userNumber, userAsBytes)
}

func (s *KYCreg) ChangeUserEmail(ctx contractapi.TransactionContextInterface, userNumber string, newEmail string) error {
	user, err := s.QueryUser(ctx, userNumber)

	if err != nil {
		return err
	}

	user.Email = newEmail

	userAsBytes, _ := json.Marshal(user)

	return ctx.GetStub().PutState(userNumber, userAsBytes)
}

// DeleteUSer deletes the user with given id in world state
func (s *KYCreg) DeleteUser(ctx contractapi.TransactionContextInterface, userNumber string) error {
	return ctx.GetStub().DelState(userNumber)
}

// UserExists bools the user with given id in world state
// func (s *KYCreg) UserExists(ctx contractapi.TransactionContextInterface, userNumber string) error {
// 	userJSON, err := ctx.GetStub().GetState(userNumber)
// 	if err != nil {
// 		return err
// 	}

// 	return userJSON != nil, nil
// }

func main() {

	chaincode, err := contractapi.NewChaincode(new(KYCreg))

	if err != nil {
		fmt.Printf("Error create KYC reg chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting KYC reg chaincode: %s", err.Error())
	}
}
