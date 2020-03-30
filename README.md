# AntiCovid19Blockchain
A Blockchain to fight the CoVid-19 spread worldwide !

## Fonctionnement
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