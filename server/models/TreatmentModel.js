const db = require('../db');
const mongoose = require('mongoose');

const TreatmentSchema = mongoose.Schema({
        treatment_id : String,
        treatment_activecomponent : String,
        treatment_description : String,
        treatment_steps : String
});

const Treatment = db.model('Treatment', TreatmentSchema);
module.exports = Treatment;