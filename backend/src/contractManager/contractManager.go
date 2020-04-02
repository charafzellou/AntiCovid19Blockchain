package contractManager

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// pub key = 0x3918EE5623bcC372282f422Cd592F564ACCd01E8
const key = `{"address":"3918ee5623bcc372282f422cd592f564accd01e8","crypto":{"cipher":"aes-128-ctr","ciphertext":"6ee348b5e048c8b104e4452763321f949333736de3a0a4e639fedd82ba9fa79a","cipherparams":{"iv":"25bb12208315d9a5116c295506b319b8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"926de8beb0bc714a755c2ce43a3bb9a16211f0ecf4e82237ba69f8285315522d"},"mac":"ea22c26fe3ed91fa0d4e7d407736e13171e9da3fd90d64874e412ecdaf2b3e57"},"id":"36595d87-880d-4b80-92a3-bf2a7f44210d","version":3}`

type contractManager struct {
	healthShare *HealthShare
}

func NewConractManager(ContractHex string, rpcUrl string) (*contractManager, error){
	conn, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	healthShare, err := NewHealthShare(common.HexToAddress(ContractHex), conn)
	if err != nil {
		return nil, err
	}
	createdContractManager := contractManager{healthShare: healthShare}
	return &createdContractManager, nil
}

func Test() (){
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	healthShare, err := NewHealthShare(common.HexToAddress("0xc2Cd0D25667CEC0f7B24Eb4a21A418d7Ef90379d"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	// Create an authorized transactor and spend 1 unicorn
	//auth, err := bind.NewTransactor(strings.NewReader(key), "finelia")
	privKey, err := crypto.HexToECDSA("98bc1a5de1d63494e638aec0c6e30a5cf62910eafdb537888ea2de8f7350ea44")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	auth := bind.NewKeyedTransactor(privKey)
	tx, err := healthShare.NewTreatment(auth, big.NewInt(1), "Milk", "This treatment work", "Step1: drink milk")
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
}

func (_contractManager *contractManager) CreatePatient(
	_id int64,
	treatmentId int64,
	screeningDate int64,
	birthYear int64,
	sex bool,
	postCode string,
	country string,
	refPhysician string,
	medicalHistory string,
	privateKey string ) (string, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	auth := bind.NewKeyedTransactor(key)
	_, err = _contractManager.healthShare.NewPatient(
		auth,
		big.NewInt(_id),
		big.NewInt(treatmentId),
		big.NewInt(birthYear),
		sex,
		postCode,
		country,
		refPhysician,
		medicalHistory,
		big.NewInt(screeningDate),
	)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf(`{ "treatmentId": %d; "screeningDate": %d; "birthYear": %d; "sexe": %t; "postCode": %s; "country": %s; "refPhysician": %s; "medicalHistory": %s}`,
		treatmentId, screeningDate, birthYear, sex, postCode, country, refPhysician, medicalHistory)
	return result, nil
}

func (_contractManager *contractManager) GetPatient(id int64, privateKey string) (string, error){
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	opts := bind.CallOpts{Pending:true, From:crypto.PubkeyToAddress(key.PublicKey)}
	treatmentId,
	screeningDate,
	remissionDate,
	deathDate, err := _contractManager.healthShare.GetPatientDetailedData(&opts, big.NewInt(id))
	if err != nil {
		return "", err
	}
	birthYear,
	sex,
	postCode,
	country,
	refPhysician,
	medicalHistory, err := _contractManager.healthShare.GetPatientPrimaryData(&opts, big.NewInt(id))
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf(`{
"treatmentId": %d; "screeningDate": %d; "remissionDate": %d; "deathDate": %d;
"birthYear": %d; "sexe": %t; "postCode": %s; "country": %s; "refPhysician": %s; "medicalHistory": %s}`,
treatmentId, screeningDate, remissionDate, deathDate, birthYear, sex, postCode, country, refPhysician, medicalHistory)
	return result, nil
}

func (_contractManager *contractManager) CreateTreatment(
	_id int64,
	activeAgent string,
	description string,
	steps string,
	privateKey string ) (string, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	auth := bind.NewKeyedTransactor(key)
	_, err = _contractManager.healthShare.NewTreatment(
		auth,
		big.NewInt(_id),
		activeAgent,
		description,
		steps,
	)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf(`{ "activeAgent": %s; "description": %s; "steps": %s}`, activeAgent, description, steps)
	return result, nil
}

func (_contractManager *contractManager) GetTreatment(id int64, privateKey string) (string, error){
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	opts := bind.CallOpts{Pending:true, From:crypto.PubkeyToAddress(key.PublicKey)}
	activeAgent,
	description,
	steps, err := _contractManager.healthShare.GetTreatment(&opts, big.NewInt(id))
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf(`{ "activeAgent": %s; "description": %s; "steps": %s }`, activeAgent, description, steps)
	return result, nil
}


