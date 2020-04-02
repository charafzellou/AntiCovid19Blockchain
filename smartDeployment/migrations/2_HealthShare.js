const HealthShare = artifacts.require("HealthShare");

module.exports = function(deployer) {
    deployer.deploy(HealthShare);
};