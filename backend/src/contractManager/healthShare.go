// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractManager

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HealthShareABI is the input ABI used to generate the binding from.
const HealthShareABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getPatientDetailedData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getPatientPrimaryData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getTreatment\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"treatmentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"birthYear\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sexe\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"postCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"country\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"refPhysician\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"medicalHistory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"screeningDate\",\"type\":\"uint256\"}],\"name\":\"newPatient\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_activeAgent\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_descrition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_steps\",\"type\":\"string\"}],\"name\":\"newTreatment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deathDate\",\"type\":\"uint256\"}],\"name\":\"patientDeath\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remissionDate\",\"type\":\"uint256\"}],\"name\":\"patientRemision\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"screeningDate\",\"type\":\"uint256\"}],\"name\":\"patientScreening\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HealthShare is an auto generated Go binding around an Ethereum contract.
type HealthShare struct {
	HealthShareCaller     // Read-only binding to the contract
	HealthShareTransactor // Write-only binding to the contract
	HealthShareFilterer   // Log filterer for contract events
}

// HealthShareCaller is an auto generated read-only Go binding around an Ethereum contract.
type HealthShareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealthShareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HealthShareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealthShareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HealthShareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealthShareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HealthShareSession struct {
	Contract     *HealthShare      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HealthShareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HealthShareCallerSession struct {
	Contract *HealthShareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HealthShareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HealthShareTransactorSession struct {
	Contract     *HealthShareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HealthShareRaw is an auto generated low-level Go binding around an Ethereum contract.
type HealthShareRaw struct {
	Contract *HealthShare // Generic contract binding to access the raw methods on
}

// HealthShareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HealthShareCallerRaw struct {
	Contract *HealthShareCaller // Generic read-only contract binding to access the raw methods on
}

// HealthShareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HealthShareTransactorRaw struct {
	Contract *HealthShareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHealthShare creates a new instance of HealthShare, bound to a specific deployed contract.
func NewHealthShare(address common.Address, backend bind.ContractBackend) (*HealthShare, error) {
	contract, err := bindHealthShare(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HealthShare{HealthShareCaller: HealthShareCaller{contract: contract}, HealthShareTransactor: HealthShareTransactor{contract: contract}, HealthShareFilterer: HealthShareFilterer{contract: contract}}, nil
}

// NewHealthShareCaller creates a new read-only instance of HealthShare, bound to a specific deployed contract.
func NewHealthShareCaller(address common.Address, caller bind.ContractCaller) (*HealthShareCaller, error) {
	contract, err := bindHealthShare(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HealthShareCaller{contract: contract}, nil
}

// NewHealthShareTransactor creates a new write-only instance of HealthShare, bound to a specific deployed contract.
func NewHealthShareTransactor(address common.Address, transactor bind.ContractTransactor) (*HealthShareTransactor, error) {
	contract, err := bindHealthShare(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HealthShareTransactor{contract: contract}, nil
}

// NewHealthShareFilterer creates a new log filterer instance of HealthShare, bound to a specific deployed contract.
func NewHealthShareFilterer(address common.Address, filterer bind.ContractFilterer) (*HealthShareFilterer, error) {
	contract, err := bindHealthShare(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HealthShareFilterer{contract: contract}, nil
}

// bindHealthShare binds a generic wrapper to an already deployed contract.
func bindHealthShare(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HealthShareABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HealthShare *HealthShareRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HealthShare.Contract.HealthShareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HealthShare *HealthShareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HealthShare.Contract.HealthShareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HealthShare *HealthShareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HealthShare.Contract.HealthShareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HealthShare *HealthShareCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HealthShare.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HealthShare *HealthShareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HealthShare.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HealthShare *HealthShareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HealthShare.Contract.contract.Transact(opts, method, params...)
}

// GetPatientDetailedData is a free data retrieval call binding the contract method 0xbc0f600b.
//
// Solidity: function getPatientDetailedData(uint256 _id) constant returns(uint256, uint256, uint256, uint256)
func (_HealthShare *HealthShareCaller) GetPatientDetailedData(opts *bind.CallOpts, _id *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _HealthShare.contract.Call(opts, out, "getPatientDetailedData", _id)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetPatientDetailedData is a free data retrieval call binding the contract method 0xbc0f600b.
//
// Solidity: function getPatientDetailedData(uint256 _id) constant returns(uint256, uint256, uint256, uint256)
func (_HealthShare *HealthShareSession) GetPatientDetailedData(_id *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _HealthShare.Contract.GetPatientDetailedData(&_HealthShare.CallOpts, _id)
}

// GetPatientDetailedData is a free data retrieval call binding the contract method 0xbc0f600b.
//
// Solidity: function getPatientDetailedData(uint256 _id) constant returns(uint256, uint256, uint256, uint256)
func (_HealthShare *HealthShareCallerSession) GetPatientDetailedData(_id *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _HealthShare.Contract.GetPatientDetailedData(&_HealthShare.CallOpts, _id)
}

// GetPatientPrimaryData is a free data retrieval call binding the contract method 0x18a95735.
//
// Solidity: function getPatientPrimaryData(uint256 _id) constant returns(uint256, bool, string, string, string, string)
func (_HealthShare *HealthShareCaller) GetPatientPrimaryData(opts *bind.CallOpts, _id *big.Int) (*big.Int, bool, string, string, string, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
		ret2 = new(string)
		ret3 = new(string)
		ret4 = new(string)
		ret5 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _HealthShare.contract.Call(opts, out, "getPatientPrimaryData", _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetPatientPrimaryData is a free data retrieval call binding the contract method 0x18a95735.
//
// Solidity: function getPatientPrimaryData(uint256 _id) constant returns(uint256, bool, string, string, string, string)
func (_HealthShare *HealthShareSession) GetPatientPrimaryData(_id *big.Int) (*big.Int, bool, string, string, string, string, error) {
	return _HealthShare.Contract.GetPatientPrimaryData(&_HealthShare.CallOpts, _id)
}

// GetPatientPrimaryData is a free data retrieval call binding the contract method 0x18a95735.
//
// Solidity: function getPatientPrimaryData(uint256 _id) constant returns(uint256, bool, string, string, string, string)
func (_HealthShare *HealthShareCallerSession) GetPatientPrimaryData(_id *big.Int) (*big.Int, bool, string, string, string, string, error) {
	return _HealthShare.Contract.GetPatientPrimaryData(&_HealthShare.CallOpts, _id)
}

// GetTreatment is a free data retrieval call binding the contract method 0xf26c5f5f.
//
// Solidity: function getTreatment(uint256 _id) constant returns(string, string, string)
func (_HealthShare *HealthShareCaller) GetTreatment(opts *bind.CallOpts, _id *big.Int) (string, string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _HealthShare.contract.Call(opts, out, "getTreatment", _id)
	return *ret0, *ret1, *ret2, err
}

// GetTreatment is a free data retrieval call binding the contract method 0xf26c5f5f.
//
// Solidity: function getTreatment(uint256 _id) constant returns(string, string, string)
func (_HealthShare *HealthShareSession) GetTreatment(_id *big.Int) (string, string, string, error) {
	return _HealthShare.Contract.GetTreatment(&_HealthShare.CallOpts, _id)
}

// GetTreatment is a free data retrieval call binding the contract method 0xf26c5f5f.
//
// Solidity: function getTreatment(uint256 _id) constant returns(string, string, string)
func (_HealthShare *HealthShareCallerSession) GetTreatment(_id *big.Int) (string, string, string, error) {
	return _HealthShare.Contract.GetTreatment(&_HealthShare.CallOpts, _id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HealthShare *HealthShareCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HealthShare.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HealthShare *HealthShareSession) Owner() (common.Address, error) {
	return _HealthShare.Contract.Owner(&_HealthShare.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HealthShare *HealthShareCallerSession) Owner() (common.Address, error) {
	return _HealthShare.Contract.Owner(&_HealthShare.CallOpts)
}

// NewPatient is a paid mutator transaction binding the contract method 0xf6805f3d.
//
// Solidity: function newPatient(uint256 _id, uint256 treatmentId, uint256 birthYear, bool sexe, string postCode, string country, string refPhysician, string medicalHistory, uint256 screeningDate) returns()
func (_HealthShare *HealthShareTransactor) NewPatient(opts *bind.TransactOpts, _id *big.Int, treatmentId *big.Int, birthYear *big.Int, sexe bool, postCode string, country string, refPhysician string, medicalHistory string, screeningDate *big.Int) (*types.Transaction, error) {
	return _HealthShare.contract.Transact(opts, "newPatient", _id, treatmentId, birthYear, sexe, postCode, country, refPhysician, medicalHistory, screeningDate)
}

// NewPatient is a paid mutator transaction binding the contract method 0xf6805f3d.
//
// Solidity: function newPatient(uint256 _id, uint256 treatmentId, uint256 birthYear, bool sexe, string postCode, string country, string refPhysician, string medicalHistory, uint256 screeningDate) returns()
func (_HealthShare *HealthShareSession) NewPatient(_id *big.Int, treatmentId *big.Int, birthYear *big.Int, sexe bool, postCode string, country string, refPhysician string, medicalHistory string, screeningDate *big.Int) (*types.Transaction, error) {
	return _HealthShare.Contract.NewPatient(&_HealthShare.TransactOpts, _id, treatmentId, birthYear, sexe, postCode, country, refPhysician, medicalHistory, screeningDate)
}

// NewPatient is a paid mutator transaction binding the contract method 0xf6805f3d.
//
// Solidity: function newPatient(uint256 _id, uint256 treatmentId, uint256 birthYear, bool sexe, string postCode, string country, string refPhysician, string medicalHistory, uint256 screeningDate) returns()
func (_HealthShare *HealthShareTransactorSession) NewPatient(_id *big.Int, treatmentId *big.Int, birthYear *big.Int, sexe bool, postCode string, country string, refPhysician string, medicalHistory string, screeningDate *big.Int) (*types.Transaction, error) {
	return _HealthShare.Contract.NewPatient(&_HealthShare.TransactOpts, _id, treatmentId, birthYear, sexe, postCode, country, refPhysician, medicalHistory, screeningDate)
}

// NewTreatment is a paid mutator transaction binding the contract method 0x03e8cdf5.
//
// Solidity: function newTreatment(uint256 _id, string _activeAgent, string _descrition, string _steps) returns()
func (_HealthShare *HealthShareTransactor) NewTreatment(opts *bind.TransactOpts, _id *big.Int, _activeAgent string, _descrition string, _steps string) (*types.Transaction, error) {
	return _HealthShare.contract.Transact(opts, "newTreatment", _id, _activeAgent, _descrition, _steps)
}

// NewTreatment is a paid mutator transaction binding the contract method 0x03e8cdf5.
//
// Solidity: function newTreatment(uint256 _id, string _activeAgent, string _descrition, string _steps) returns()
func (_HealthShare *HealthShareSession) NewTreatment(_id *big.Int, _activeAgent string, _descrition string, _steps string) (*types.Transaction, error) {
	return _HealthShare.Contract.NewTreatment(&_HealthShare.TransactOpts, _id, _activeAgent, _descrition, _steps)
}

// NewTreatment is a paid mutator transaction binding the contract method 0x03e8cdf5.
//
// Solidity: function newTreatment(uint256 _id, string _activeAgent, string _descrition, string _steps) returns()
func (_HealthShare *HealthShareTransactorSession) NewTreatment(_id *big.Int, _activeAgent string, _descrition string, _steps string) (*types.Transaction, error) {
	return _HealthShare.Contract.NewTreatment(&_HealthShare.TransactOpts, _id, _activeAgent, _descrition, _steps)
}

// PatientDeath is a paid mutator transaction binding the contract method 0x2c4bef34.
//
// Solidity: function patientDeath(uint256 _id, uint256 deathDate) returns()
func (_HealthShare *HealthShareTransactor) PatientDeath(opts *bind.TransactOpts, _id *big.Int, deathDate *big.Int) (*types.Transaction, error) {
	return _HealthShare.contract.Transact(opts, "patientDeath", _id, deathDate)
}

// PatientDeath is a paid mutator transaction binding the contract method 0x2c4bef34.
//
// Solidity: function patientDeath(uint256 _id, uint256 deathDate) returns()
func (_HealthShare *HealthShareSession) PatientDeath(_id *big.Int, deathDate *big.Int) (*types.Transaction, error) {
	return _HealthShare.Contract.PatientDeath(&_HealthShare.TransactOpts, _id, deathDate)
}

// PatientDeath is a paid mutator transaction binding the contract method 0x2c4bef34.
//
// Solidity: function patientDeath(uint256 _id, uint256 deathDate) returns()
func (_HealthShare *HealthShareTransactorSession) PatientDeath(_id *big.Int, deathDate *big.Int) (*types.Transaction, error) {
	return _HealthShare.Contract.PatientDeath(&_HealthShare.TransactOpts, _id, deathDate)
}

// PatientRemision is a paid mutator transaction binding the contract method 0xd5b4724c.
//
// Solidity: function patientRemision(uint256 _id, uint256 remissionDate) returns()
func (_HealthShare *HealthShareTransactor) PatientRemision(opts *bind.TransactOpts, _id *big.Int, remissionDate *big.Int) (*types.Transaction, error) {
	return _HealthShare.contract.Transact(opts, "patientRemision", _id, remissionDate)
}

// PatientRemision is a paid mutator transaction binding the contract method 0xd5b4724c.
//
// Solidity: function patientRemision(uint256 _id, uint256 remissionDate) returns()
func (_HealthShare *HealthShareSession) PatientRemision(_id *big.Int, remissionDate *big.Int) (*types.Transaction, error) {
	return _HealthShare.Contract.PatientRemision(&_HealthShare.TransactOpts, _id, remissionDate)
}

// PatientRemision is a paid mutator transaction binding the contract method 0xd5b4724c.
//
// Solidity: function patientRemision(uint256 _id, uint256 remissionDate) returns()
func (_HealthShare *HealthShareTransactorSession) PatientRemision(_id *big.Int, remissionDate *big.Int) (*types.Transaction, error) {
	return _HealthShare.Contract.PatientRemision(&_HealthShare.TransactOpts, _id, remissionDate)
}

// PatientScreening is a paid mutator transaction binding the contract method 0x1c5dff0f.
//
// Solidity: function patientScreening(uint256 _id, uint256 screeningDate) returns()
func (_HealthShare *HealthShareTransactor) PatientScreening(opts *bind.TransactOpts, _id *big.Int, screeningDate *big.Int) (*types.Transaction, error) {
	return _HealthShare.contract.Transact(opts, "patientScreening", _id, screeningDate)
}

// PatientScreening is a paid mutator transaction binding the contract method 0x1c5dff0f.
//
// Solidity: function patientScreening(uint256 _id, uint256 screeningDate) returns()
func (_HealthShare *HealthShareSession) PatientScreening(_id *big.Int, screeningDate *big.Int) (*types.Transaction, error) {
	return _HealthShare.Contract.PatientScreening(&_HealthShare.TransactOpts, _id, screeningDate)
}

// PatientScreening is a paid mutator transaction binding the contract method 0x1c5dff0f.
//
// Solidity: function patientScreening(uint256 _id, uint256 screeningDate) returns()
func (_HealthShare *HealthShareTransactorSession) PatientScreening(_id *big.Int, screeningDate *big.Int) (*types.Transaction, error) {
	return _HealthShare.Contract.PatientScreening(&_HealthShare.TransactOpts, _id, screeningDate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_HealthShare *HealthShareTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _HealthShare.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_HealthShare *HealthShareSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _HealthShare.Contract.TransferOwnership(&_HealthShare.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_HealthShare *HealthShareTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _HealthShare.Contract.TransferOwnership(&_HealthShare.TransactOpts, _newOwner)
}

// HealthShareOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HealthShare contract.
type HealthShareOwnershipTransferredIterator struct {
	Event *HealthShareOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HealthShareOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealthShareOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HealthShareOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HealthShareOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealthShareOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealthShareOwnershipTransferred represents a OwnershipTransferred event raised by the HealthShare contract.
type HealthShareOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HealthShare *HealthShareFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HealthShareOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HealthShare.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HealthShareOwnershipTransferredIterator{contract: _HealthShare.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HealthShare *HealthShareFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HealthShareOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HealthShare.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealthShareOwnershipTransferred)
				if err := _HealthShare.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HealthShare *HealthShareFilterer) ParseOwnershipTransferred(log types.Log) (*HealthShareOwnershipTransferred, error) {
	event := new(HealthShareOwnershipTransferred)
	if err := _HealthShare.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}