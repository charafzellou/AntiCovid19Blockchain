# AntiCovid19Blockchain
A Blockchain to fight the CoVid-19 spread worldwide !

## Basic Database Modelization
```
Frontend:
    - Gestion d'entrée de données

Solidity:
	- Patient :
			ID PATIENT
			ID TRAITEMENT
			Date de Naissance
			Sexe
			Code Postal
			Pays
			Médecin Traitant
			Antécédents
			Date de Dépistage
			Date de Rémission
			Date de Décès
	- Traitement :
			ID TRAITEMENT
			Composant actif
			Description
			Etapes du traitement

GO :
    - Liaison entre Front et Solidity
    - Liaison entre Front et MongoDB

MongoDB:
	- Patient : 
			ID PATIENT
			* Nom
			* Prénom
			* Numéro SS
			* Numéro Téléphonique
			* Adresse
	- Admission :
			ID PATIENT
			ID TRAITEMENT
			Date de Naissance
			Sexe
			Code Postal
			Pays
			Médecin Traitant
			Antécédents
			Date de Dépistage
			Date de Rémission
			Date de Décès
	- Traitement :
			ID TRAITEMENT
			Composant actif
			Description
			Etapes du traitement
```

## Launching the app
To launch the app, use the following commands:
```bash
git clone https://github.com/charafzellou/AntiCovid19Blockchain
cd AntiCovid19Blockchain
docker-compose up -d
```
## Ports used by app
Ports used by Docker-Compose:
```
Mongo --> 27017:27017
Mongo-Express GUI --> 8081:8081
API React for Website --> 3000:3000
React Client for Website --> 8686:8686
Ganache-Cli for Ethereum Smart Contracts --> 8545:8545
```

## Fixing issues related to NPM
For reasons unknown to my modest brain and knowledge, sometimes NPM refuses to cooperate properly with Docker to initialize the Node Modules.

In order to bypass this issue, you need to initialize them before stting up the Docker environment, run this after cloning the Git:

```bash
docker rm $(docker ps -aq) && docker-compose down
cd api
npm install
cd ../client/src/
npm install
cd ../../server/
npm install
```

## Scrap the whole environment
If you face an undocumented issue and want to scrap everything to restart, it is recommended to get to the top folder `./AntiCovid19Blockchain/` and run this:
```bash
docker rm -f $(docker ps -aq) && docker-compose down
```