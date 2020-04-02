module.exports = {
    networks: {
        development: {
            host: "ganachecli://ganache-cli",     // Localhost (default: none)
            port: 8545,            // Standard Ethereum port (default: none)
            network_id: "5777",       // Any network (default: none)
        },
    },
    mocha: {
        // timeout: 100000
    },
    compilers: {
        solc: {
            // version: "0.5.1",    // Fetch exact version from solc-bin (default: truffle's version)
            // docker: true,        // Use "0.5.1" you've installed locally with docker (default: false)
            // settings: {          // See the solidity docs for advice about optimization and evmVersion
            //  optimizer: {
            //    enabled: false,
            //    runs: 200
            //  },
            //  evmVersion: "byzantium"
            // }
        }
    }
}