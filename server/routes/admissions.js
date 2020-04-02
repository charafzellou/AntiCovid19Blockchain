const express = require('express');
const router = express.Router();
const AdmissionModel = require('../models/AdmissionModel')

router.get('/', (request, response) => {
	AdmissionModel.find().then(data => response.json(data));
});

router.get('/:id', (request, response) => {
	AdmissionModel.findOne({patient_id: request.params.id}).then(data => response.json(data));
});

router.post('/', (request, response) => {
	const admission = new AdmissionModel(request.body);
	console.log('################');
	console.log(admission);
	console.log('################');
	admission.save(request.body)
		.then(admission => res.status(201).json(admission))
		.catch(error => { if(error.name === "ValidationError") {
			response.status(400).json(error.errors);
			} else {
				response.sendStatus(500);
			}
	});
});

router.delete('/:id', (request, response) => {
	AdmissionModel.deleteOne({patient_id: request.params.id})
	.then(() => response.sendStatus(204));
});

router.put('/:id', (request, response) => {
	AdmissionModel.findByIdAndUpdate(
		{patient_id: request.params.id}, request.body)
		.then(data => response.json(...data, ...req.body));
});

module.exports = router;
