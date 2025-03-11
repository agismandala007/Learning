const path = require("path");
const fs = require("fs");
const solc = require("solc");
const { interfaces } = require("mocha");

const inboxPath = path.resolve(__dirname, "contracts", "Inbox.sol");
const source = fs.readFileSync(inboxPath, "utf-8");

const input = {
  language: "Solidity",
  sources: {
    "Inbox.sol": {
      content: source,
    },
  },
  settings: {
    outputSelection: {
      "*": {
        "*": ["*"],
      },
    },
  },
};

const output = JSON.parse(solc.compile(JSON.stringify(input)));

if (output.errors) {
  output.errors.forEach((error) => {
    console.log(error.formattedMessage);
  });

  throw new Error("Compilation failed.");
}

const inboxContract = output.contracts["Inbox.sol"]["Inbox"];

if (!inboxContract) {
  throw new Error("Contract 'Inbox' not found");
}

module.exports = {
  interfaces: JSON.stringify(inboxContract.abi),
  bytecode: inboxContract.evm.bytecode.object,
};
