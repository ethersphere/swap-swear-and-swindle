const { ethers, JsonRpcProvider } = require("ethers");
const fs = require('fs');

// Replace with your RPC URL and contract details
const rpcUrl = "https://eth-goerli.g.alchemy.com/v2/UbKbbJpAxim12srTJ5vUo3DdQy0WPdHK";
const provider = new JsonRpcProvider('https://go.getblock.io/3ce184cf4bf44236b537fdb3b6d53a29');
const contractAddress = "0x73c412512E1cA0be3b89b77aB3466dA6A1B9d273";
const contractABI = [
  {
    inputs: [
      { internalType: "address", name: "_ERC20Address", type: "address" },
    ],
    stateMutability: "nonpayable",
    type: "constructor",
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: "address",
        name: "contractAddress",
        type: "address",
      },
    ],
    name: "SimpleSwapDeployed",
    type: "event",
  },
  {
    inputs: [],
    name: "ERC20Address",
    outputs: [{ internalType: "address", name: "", type: "address" }],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      { internalType: "address", name: "issuer", type: "address" },
      {
        internalType: "uint256",
        name: "defaultHardDepositTimeoutDuration",
        type: "uint256",
      },
      { internalType: "bytes32", name: "salt", type: "bytes32" },
    ],
    name: "deploySimpleSwap",
    outputs: [{ internalType: "address", name: "", type: "address" }],
    stateMutability: "nonpayable",
    type: "function",
  },
  {
    inputs: [{ internalType: "address", name: "", type: "address" }],
    name: "deployedContracts",
    outputs: [{ internalType: "bool", name: "", type: "bool" }],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "master",
    outputs: [{ internalType: "address", name: "", type: "address" }],
    stateMutability: "view",
    type: "function",
  },
];

// Specify the starting block
const startBlock = 10012821; // Replace with the block number from where you want to start

// Create a contract instance
const myContract = new ethers.Contract(contractAddress, contractABI, provider);

// Function to decode transaction input data
async function decodeTransactionInput(transactionHash) {
  try {
    // Get the transaction
    const transaction = await provider.getTransaction(transactionHash);
    if (!transaction) {
      console.log("Transaction not found!");
      return null;
    }

    // Create an interface from the ABI to decode the data
    const contractInterface = new ethers.Interface(contractABI);

    // Decode the transaction input data
    const decodedInput = contractInterface.parseTransaction({ data: transaction.data });

    // Convert decoded data to a more readable format
    return {
      functionName: decodedInput.name,
      args: decodedInput.args
    };
  } catch (error) {
    console.error("Error decoding transaction input:", error);
    return null;
  }
}


// Prepare to store the decoded data
let decodedDataArray = [];

// Query past events and save decoded data
async function getPastEvents() {
  const events = await myContract.queryFilter("SimpleSwapDeployed", startBlock, "latest");
  for (let event of events) {
    const decodedData = await decodeTransactionInput(event.transactionHash);
    if (decodedData) {
      decodedDataArray.push(decodedData);
    }
  }

  // Custom replacer function to handle BigInt serialization
  function replacer(key, value) {
    if (typeof value === 'bigint') {
      return value.toString();
    }
    return value;
  }

  // Save the decoded data array to a JSON file
  fs.writeFile('migration/decodedData.json', JSON.stringify(decodedDataArray, replacer, 2), (err) => {
    if (err) {
      console.error('Error writing file:', err);
    } else {
      console.log('Successfully wrote decoded data to decodedData.json');
    }
  });
}

// Call
// Call the function to start the process
getPastEvents().catch(console.error);