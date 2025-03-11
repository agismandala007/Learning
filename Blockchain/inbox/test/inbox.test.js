const assert = require("assert");
const ganache = require("ganache");
const { beforeEach } = require("mocha");
const { Web3 } = require("web3");
const { interface, bytecode } = require("../compile");

const web3 = new Web3(ganache.provider());

let accounts;
let inbox;

beforeEach(async () => {
  // Get a list of all accounts
  accounts = await web3.eth.getAccounts();

  inbox = await new web3.eth.Contract(JSON.parse(interface))
    .deploy({ data: bytecode, arguments: ["Hi There"] })
    .send({ from: accounts[0], gas: "1000000" });

  // Use one of those accounts to deploy
  //the contract
});

describe("Inbox", () => {
  it("deploy a contract", () => {
    console.log(inbox);
  });
});
