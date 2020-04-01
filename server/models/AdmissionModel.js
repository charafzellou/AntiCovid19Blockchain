const db = require('../db');
const mongoose = require('mongoose');

const AdmissionSchema = mongoose.Schema({
        patient_id : String,
        treatment_name : String,
        admission_birthdate : String,
        admission_gender : String,
        admission_postalcode : String,
        admission_country : String,
        admission_assigneddoctor : String,
        admission_knowndiseases : String,
        admission_dateof_screening : String,
        admission_dateof_remission : String,
        admission_dateof_death : String
});

const Admission = db.model('Admission', AdmissionSchema);
module.exports = Admission;