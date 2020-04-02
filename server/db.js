const mongoose = require('mongoose');
const mongoHost = 'mongodb://mongodb:27017/';

mongoose.connect(mongoHost, {
	user: "root",
	pass: "password",
	dbName: "anticovid",
	useNewUrlParser: true
});

var db = mongoose.connection;
db.on('error', console.error.bind(console, 'Connection error:'));
db.once('open', function() {
  console.log('Connection to MongoDB successful!');
});

module.exports = mongoose.connection;