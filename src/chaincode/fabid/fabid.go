/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	// "github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// "strconv"

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

// Car describes basic details of what makes up a car
type Identication struct {
	ObjectType string `json:"docType"`
	Id   string `json:"id"`
	Name   string `json:"name"`
	DateOfbirth  string `json:"dateOfbirth"`
	Sex string `json:"sex"`
	Nationality  string `json:"nationality"`
	PlaceOforigin  string `json:"placeOforigin"`
	PlaceOfresidence  string `json:"placeOfresidence"`
	DateOfExpiry  string `json:"dateOfExpiry"`
	PersonalIdentification  string `json:"personalIdentification"`
	CreatedDate  string `json:"createdDate"`
	PersonSignature  string `json:"personSignature"`
	AvatarUrl  string `json:"avatarUrl"`
	SignatureUrl  string `json:"signatureUrl"`
	IsDelete string `json:"isDelete"`
}

type Admin struct {
	ObjectType string `json:"docType"`
	Id string `json:"id"`
	Email string `json:"email"`
	Password string `json:"Password"`
}

// HistoryQueryResult structure used for returning result of history query
type HistoryQueryResult struct {
	Record    *Identication    `json:"record"`
	TxId      string    `json:"txId"`
	Timestamp time.Time `json:"timestamp"`
	IsDelete  bool      `json:"isDelete"`
}

// QueryResult structure used for handling result of query
// type QueryResult struct {
// 	Key    string `json:"Key"`
// 	Record *Identication
// }

// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	identications := []Identication{
		// Identication{ObjectType: "identication", Id: "048201000569", Name: "Võ Thành Vĩ", DateOfbirth: "24-08-2001", Sex: "Nam", Nationality: "Việt Nam", PlaceOforigin: "Duy Châu, Duy Xuyên, Quảng Nam", PlaceOfresidence: "Tổ 23 Mân Thái, Sơn Trà, Đà Nẵng", DateOfExpiry: "24-08-2026", PersonalIdentification: "Sẹo chấm C 2cm sau cánh mũi", CreatedDate: "29-04-2021", PersonSignature: "CỤC TRƯỞNG CỤC CẢNH SÁT QUẢN Lý HÀNH CHÍNH VỀ TRẬT TỰ XÃ HỘI, Phạm Công Nguyên", AvatarUrl: "https://res.cloudinary.com/h-b-ch-khoa/image/upload/v1670868317/i5ew932zswwmqtaynpyl.jpg", SignatureUrl: "https://res.cloudinary.com/h-b-ch-khoa/image/upload/v1670868259/fxpuj5iwyepl0qu7qma7.jpg", IsDelete: "false"},
		Identication{ObjectType: "identication", Id: "048201000570", Name: "Huỳnh Tấn Phát", DateOfbirth: "24-02-2001", Sex: "Nam", Nationality: "Việt Nam", PlaceOforigin: "Duy Châu, Duy Xuyên, Quảng Nam", PlaceOfresidence: "Tổ 35 Mân Thái, Sơn Trà, Đà Nẵng", DateOfExpiry: "24-08-2026", PersonalIdentification: "Sẹo chấm C 10cm sau cánh mũi", CreatedDate: "29-04-2021", PersonSignature: "CỤC TRƯỞNG CỤC CẢNH SÁT QUẢN Lý HÀNH CHÍNH VỀ TRẬT TỰ XÃ HỘI, Phạm Công Nguyên", AvatarUrl: "https://res.cloudinary.com/h-b-ch-khoa/image/upload/v1670871091/vvjkyncs9s5gvntcrwux.jpg", SignatureUrl: "https://res.cloudinary.com/h-b-ch-khoa/image/upload/v1670868259/fxpuj5iwyepl0qu7qma7.jpg", IsDelete: "true"},
		Identication{ObjectType: "identication", Id: "048201000571", Name: "Nguyễn Tuấn Hùng", DateOfbirth: "15-08-2001", Sex: "Nam", Nationality: "Việt Nam", PlaceOforigin: "Duy Châu, Duy Xuyên, Quảng Nam", PlaceOfresidence: "Tổ 23 Mân Thái, Sơn Trà, Đà Nẵng", DateOfExpiry: "24-08-2026", PersonalIdentification: "Sẹo chấm C 2cm sau cánh mũi", CreatedDate: "29-04-2021", PersonSignature: "CỤC TRƯỞNG CỤC CẢNH SÁT QUẢN Lý HÀNH CHÍNH VỀ TRẬT TỰ XÃ HỘI, Phạm Công Nguyên", AvatarUrl: "https://res.cloudinary.com/h-b-ch-khoa/image/upload/v1670871666/a9lpqm5jvacuu3erakqe.jpg", SignatureUrl: "https://res.cloudinary.com/h-b-ch-khoa/image/upload/v1670868259/fxpuj5iwyepl0qu7qma7.jpg", IsDelete: "true"},
	}

	admins := []Admin{
		Admin{ObjectType: "admin", Id: "123", Email: "admin@gmail.com", Password: "adminpw"},
	}

	for _, identication := range identications {
		identicationAsBytes, _ := json.Marshal(identication)
		err := ctx.GetStub().PutState("IDENTICATION" + identication.Id, identicationAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	for _, admin := range admins {
		adminAsBytes, _ := json.Marshal(admin)
		err := ctx.GetStub().PutState("ADMIN" + admin.Id, adminAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

// CreateCar adds a new car to the world state with given details
func (s *SmartContract) CreateIdentication(ctx contractapi.TransactionContextInterface, id string, name string, dateOfbirth string, sex string, nationality string, placeOforigin string, placeOfresidence string, dateOfExpiry string, personalIdentification string, personSignature string, avatarUrl string, signatureUrl string, createdDate string, isDelete string) error {
	identication := Identication{
		ObjectType : "identication",
		Id : id,
		Name : name,
		DateOfbirth : dateOfbirth,
		Sex : sex,
		Nationality : nationality,
		PlaceOforigin : placeOforigin,
		PlaceOfresidence : placeOfresidence,
		DateOfExpiry : dateOfExpiry,
		PersonalIdentification : personalIdentification,
		PersonSignature : personSignature,
		AvatarUrl : avatarUrl,
		SignatureUrl : signatureUrl,
		CreatedDate : createdDate,
		IsDelete : isDelete,
	}

	identicationAsBytes, _ := json.Marshal(identication)

	return ctx.GetStub().PutState("IDENTICATION" + id, identicationAsBytes)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QueryIdentication(ctx contractapi.TransactionContextInterface, id string) (*Identication, error) {
	identicationAsBytes, err := ctx.GetStub().GetState("IDENTICATION" + id)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if identicationAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", id)
	}

	identication := new(Identication)
	_ = json.Unmarshal(identicationAsBytes, identication)

	return identication, nil
}

// QueryAllCars returns all cars found in world state
func (s *SmartContract) QueryAllIdentications(ctx contractapi.TransactionContextInterface) ([]Identication, error) {
	startKey := "IDENTICATION"
	endKey := "IDENTICATIONZ"

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []Identication{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		identication := new(Identication)
		_ = json.Unmarshal(queryResponse.Value, identication)

		queryResult := Identication{ObjectType: identication.ObjectType, Id: identication.Id, Name: identication.Name, DateOfbirth: identication.DateOfbirth, Sex: identication.Sex, Nationality: identication.Nationality, PlaceOforigin: identication.PlaceOforigin, PlaceOfresidence: identication.PlaceOfresidence, DateOfExpiry: identication.DateOfExpiry, PersonalIdentification: identication.PersonalIdentification, CreatedDate: identication.CreatedDate, PersonSignature: identication.PersonSignature, AvatarUrl: identication.AvatarUrl, SignatureUrl: identication.SignatureUrl, IsDelete: identication.IsDelete}
		results = append(results, queryResult)
	}

	return results, nil
}

// ChangeCarOwner updates the owner field of car with given id in world state
func (s *SmartContract) UpdateIdentication(ctx contractapi.TransactionContextInterface, id string, name string, dateOfbirth string, sex string, nationality string, placeOforigin string, placeOfresidence string, dateOfExpiry string, personalIdentification string, personSignature string, avatarUrl string, signatureUrl string, createdDate string, isDelete string) error {
	identication, err := s.QueryIdentication(ctx, id)

	if err != nil {
		return err
	}

	identication.Name = name
	identication.DateOfbirth = dateOfbirth
	identication.Sex = sex
	identication.Nationality = nationality
	identication.PlaceOforigin = placeOforigin
	identication.PlaceOfresidence = placeOfresidence
	identication.DateOfExpiry = dateOfExpiry
	identication.PersonalIdentification = personalIdentification
	identication.PersonSignature = personSignature
	identication.AvatarUrl = avatarUrl
	identication.SignatureUrl = signatureUrl
	identication.CreatedDate = createdDate
	identication.IsDelete = isDelete
	identicationAsBytes, _ := json.Marshal(identication)

	return ctx.GetStub().PutState("IDENTICATION" + id, identicationAsBytes)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteIdentication(ctx contractapi.TransactionContextInterface, id string) error {
	return s.DeleteAsset(ctx, "IDENTICATION" + id)
}

// ADMIN 
func (s *SmartContract) QueryAdmin(ctx contractapi.TransactionContextInterface, id string) (*Admin, error) {
	adminAsBytes, err := ctx.GetStub().GetState("ADMIN" + id)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if adminAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", id)
	}

	admin := new(Admin)
	_ = json.Unmarshal(adminAsBytes, admin)

	return admin, nil
}


// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// GetAssetHistory returns the chain of custody for an asset since issuance.
func (t *SmartContract) GetAssetHistory(ctx contractapi.TransactionContextInterface, assetID string) ([]HistoryQueryResult, error) {
	log.Printf("GetAssetHistory: ID %v", assetID)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(assetID)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Identication
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &asset)
			if err != nil {
				return nil, err
			}
		} else {
			asset = Identication{
				Id: assetID,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: timestamp,
			Record:    &asset,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}	

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
