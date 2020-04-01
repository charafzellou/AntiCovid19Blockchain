const express = require('express');
const bodyparser = require('body-parser');
const admissionRouter = require('./routes/admissions');
const patientRouter = require('./routes/patients');
const treatmentRouter = require('./routes/treatments');
const app = express();

app.use(bodyparser.json());

app.use((request, response, next) => {
	console.log(request.headers);
	next();
	}
);

app.use('/admission', admissionRouter);
app.use('/patient', patientRouter);
app.use('/treatment', treatmentRouter);

app.listen(3000, () => console.log('Listening on 3000: ' + Date.now()));
