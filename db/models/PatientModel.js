const db = require('../db');
const mongoose = require('mongoose');

const PatientSchema = mongoose.Schema({
        patient_id : String,
        patient_name: String,
        patient_surname: String,
        patient_socialnumber: String,
        patient_phonenumber: String,
        patient_address: String
});

const Patient = db.model('Patient', PatientSchema);
module.exports = Patient;