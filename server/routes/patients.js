const express = require('express');
const router = express.Router();
const PatientModel = require('../models/PatientModel')

router.get('/', (request, response) => {
	PatientModel.find().then(data => response.json(data));
});

router.get('/:id', (request, response) => {
	PatientModel.findOne({coin_id: request.params.id}).then(data => response.json(data));
});

router.post('/', (request, response) => {
	const patient = new PatientModel(request.body);
	console.log('################');
	console.log(patient);
	console.log('################');
	patient.save(request.body)
		.then(patient => res.status(201).json(patient))
		.catch(error => { if(error.name === "ValidationError") {
			response.status(400).json(error.errors);
			} else {
				response.sendStatus(500);
			}
	});
});

router.delete('/:id', (request, response) => {
	PatientModel.deleteOne({coin_id: request.params.id})
	.then(() => response.sendStatus(204));
});

router.put('/:id', (request, response) => {
	PatientModel.findByIdAndUpdate(
		{coin_id: request.params.id}, request.body)
		.then(data => response.json(...data, ...req.body));
});

module.exports = router;
