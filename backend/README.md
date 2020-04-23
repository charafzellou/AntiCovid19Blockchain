# Backend

## Get Started

To build the app
```Shell
cd src/main
go build
```

to start the app

```Shell
./main
```

2 flags are available 

- contractAddr string: The address of the deployed HealthContract
- rpcUrl string : RPC url (default "http://localhost:8545")

```Shell
./main -contractAddr=0x0000000000000000000000000000000000000000 -rpcUrl=http://localhost:8545
```

In order to work properly these flags are mandatory

## Endpoints

### GET

/v1/treatments/{id}

/v1/patients/{id}

### POST

create treatments :

/v1/treatments

Params :
- id : int
- activeAgent : string
- description : string
- steps : string
- jsonKey : string
- passphrase : string

Response : 
```json
{ 
    "activeAgent": "string",
    "description": "string",
    "steps":       "string"
}
```

/v1/patients

Params :
- id : int
- treatmentId: int;
- screeningDate: int;
- birthYear: int;
- sex: bool; 
- postCode: string;
- country: string;
- refPhysician: string;
- medicalHistory: string
- jsonKey : string
- passphrase : string

Response :
```json
{ 
    "treatmentId":    "int",
    "screeningDate":  "int",
    "birthYear":      "int",
    "sex":           "bool", 
    "postCode":       "string",
    "country":        "string",
    "refPhysician":   "string",
    "medicalHistory": "string"
}
````


/v1/patients/{id}/screening

Params :
- screening : int
- jsonKey : string
- passphrase : string

Response :
```json
{ 
    "id": "int",
    "success": true
}
```

/v1/patients/{id}/death

Params :
- death : int
- jsonKey : string
- passphrase : string

Response :
```json
{ 
    "id": "int",
    "success": true
}
```

/v1/patients/{id}/remission

Params :
- remission : int
- jsonKey : string
- passphrase : string

Response :
```json
{ 
    "id": "int",
    "success": true
}
```

### Errors

When an error occur the response is along with error 500:
```json
{ 
    "error": "description"
}
```

