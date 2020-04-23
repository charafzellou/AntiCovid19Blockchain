module main

replace anticovid/routes => ../routes

replace anticovid/contractManager => ../contractManager

go 1.12

require (
	anticovid/contractManager v0.0.0-00010101000000-000000000000 // indirect
	anticovid/routes v0.0.0-00010101000000-000000000000
	github.com/ethereum/go-ethereum v1.9.12 // indirect
	github.com/gorilla/mux v1.7.4
)
