const express = require('express');
const router = express.Router();
const TreatmentModel = require('../models/TreatmentModel')

router.get('/', (request, response) => {
	TreatmentModel.find().then(data => response.json(data));
});

router.get('/:id', (request, response) => {
	TreatmentModel.findOne({treatment_id: request.params.id}).then(data => response.json(data));
});

router.post('/', (request, response) => {
	const treatment = new TreatmentModel(request.body);
	console.log('################');
	console.log(treatment);
	console.log('################');
	treatment.save(request.body)
		.then(treatment => res.status(201).json(treatment))
		.catch(error => { if(error.name === "ValidationError") {
			response.status(400).json(error.errors);
			} else {
				response.sendStatus(500);
			}
	});
});

router.delete('/:id', (request, response) => {
	TreatmentModel.deleteOne({treatment_id: request.params.id})
	.then(() => response.sendStatus(204));
});

router.put('/:id', (request, response) => {
	TreatmentModel.findByIdAndUpdate(
		{treatment_id: request.params.id}, request.body)
		.then(data => response.json(...data, ...req.body));
});

module.exports = router;
