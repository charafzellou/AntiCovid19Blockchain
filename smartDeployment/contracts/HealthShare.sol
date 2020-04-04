pragma solidity ^0.5.0;
import "./ownable.sol";

contract HealthShare is Ownable {
    struct Treatment {
        address payable treatmentOwner;
        string activeAgent;
        string description;
        string steps;
    }

    struct Patient {
        address payable dataOwner;
        uint treatmentId;
        uint birthYear;
        bool sexe;
        string postCode;
        string country;
        string refPhysician;
        string medicalHistory;
        uint screeningDate;
        uint remissionDate;
        uint deathDate;
    }

    mapping(uint => Treatment) treatments;
    mapping(uint => Patient) patients;

    modifier onlyDataOwner(uint _id) {
        require(msg.sender == patients[_id].dataOwner, "you dont own this data");
        _;
    }

    modifier onlyTreatmentOwner(uint _id) {
        require(msg.sender == treatments[_id].treatmentOwner, "you dont own this treatment");
        _;
    }

    modifier issetPatient(uint _id) {
        require(patientExist(_id), "this patientId doesn't exist");
        _;
    }

    modifier issetTreatment(uint _id) {
        require(treatmentExist(_id), "this TreatmentId doesn't exist");
        _;
    }

    function newTreatment(
        uint _id,
        string calldata _activeAgent,
        string calldata _descrition,
        string calldata _steps
    ) external payable{
        if ( treatmentExist(_id) ) {
            require(false, "this treatmentId already exist");
        }
        Treatment storage treatment = treatments[_id];
        treatment.treatmentOwner = msg.sender;
        treatment.activeAgent = _activeAgent;
        treatment.description = _descrition;
        treatment.steps = _steps;
    }

    function newPatient(
        uint _id,
        uint treatmentId,
        uint birthYear,
        bool sexe,
        string calldata postCode,
        string calldata country,
        string calldata refPhysician,
        string calldata medicalHistory,
        uint screeningDate
    ) external payable{
        if ( !treatmentExist(treatmentId) ) {
            require(false, "this treatmentId doesn't exist");
        }
        if ( patientExist(_id) ){
            require(false, "this patientId already exist");
        }
        Patient storage patient = patients[_id];
        patient.dataOwner = msg.sender;
        patient.treatmentId = treatmentId;
        patient.birthYear = birthYear;
        patient.sexe = sexe;
        patient.postCode = postCode;
        patient.country = country;
        patient.refPhysician = refPhysician;
        patient.medicalHistory = medicalHistory;
        patient.screeningDate = screeningDate;
    }

    function patientDeath(
        uint _id,
        uint deathDate
    ) external onlyDataOwner(_id) issetPatient(_id) payable {
        patients[_id].deathDate = deathDate;
    }

    function patientRemision(
        uint _id,
        uint remissionDate
    ) external onlyDataOwner(_id) issetPatient(_id) payable {
        patients[_id].remissionDate = remissionDate;
    }

    function patientScreening(
        uint _id,
        uint screeningDate
    ) external onlyDataOwner(_id) issetPatient(_id) payable {
        patients[_id].screeningDate = screeningDate;
    }

    function getTreatment(
        uint _id
    ) external issetTreatment(_id) view returns (
        string memory,
        string memory,
        string memory
    ) {
        return (
            treatments[_id].activeAgent,
            treatments[_id].description,
            treatments[_id].steps
        );
    }

    function getPatientPrimaryData(
        uint _id
    ) external issetPatient(_id) view returns (
        uint,
        bool,
        string memory,
        string memory,
        string memory,
        string memory
    ){
        return (
            patients[_id].birthYear,
            patients[_id].sexe,
            patients[_id].postCode,
            patients[_id].country,
            patients[_id].refPhysician,
            patients[_id].medicalHistory
        );
    }

    function getPatientDetailedData(
        uint _id
    ) external issetPatient(_id) view returns (
        uint,
        uint,
        uint,
        uint
    ){
        return (
            patients[_id].treatmentId,
            patients[_id].screeningDate,
            patients[_id].remissionDate,
            patients[_id].deathDate
        );
    }

    function treatmentExist(
        uint id
    ) private view returns (bool){
        return treatments[id].treatmentOwner != address(0);
    }

    function patientExist(
        uint id
    ) private view returns (bool){
        return patients[id].dataOwner != address(0);
    }
}