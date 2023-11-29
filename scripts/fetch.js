const hre = require("hardhat");
const fs = require("fs/promises");

// Function to decode transaction input data
async function decodeTransactionInput(transactionHash, provider) {
  try {
    // Get the transaction

    const transaction = await provider.getTransaction(transactionHash);
    if (!transaction) {
      console.log("Transaction not found!");
      return null;
    }
    const contractABI =
      await require("../artifacts/contracts/SimpleSwapFactory.sol/SimpleSwapFactory.json");
    // Create an interface from the ABI to decode the data
    const contractInterface = new ethers.Interface(contractABI.abi);

    // Decode the transaction input data
    const decodedInput = contractInterface.parseTransaction({
      data: transaction.data,
    });

    // Convert decoded data to a more readable format
    return {
      functionName: decodedInput.name,
      args: decodedInput.args,
    };
  } catch (error) {
    console.error("Error decoding transaction input:", error);
    return null;
  }
}
// Custom replacer function to handle BigInt serialization
function replacer(key, value) {
  if (typeof value === "bigint") {
    return value.toString();
  }
  return value;
}

async function main() {
  const provider = ethers.provider;

  // Specify the starting block
  const startBlock = 10012821; // Replace with the block number from where you want to start

  const contractAddress = "0x73c412512E1cA0be3b89b77aB3466dA6A1B9d273";
  const myContract = await ethers.getContractAt(
    "SimpleSwapFactory",
    contractAddress
  );

  // Prepare to store the decoded data
  let decodedDataArray = [];

  // Define batch size and delay (in milliseconds)
  const batchSize = 10000; // Adjust based on your needs
  const delay = 20; // Delay of 0.02 seconds

  // Get the latest block number
  const latestBlock = await provider.getBlockNumber();

  for (
    let currentBlock = startBlock;
    currentBlock < latestBlock;
    currentBlock += batchSize
  ) {
    // Calculate the end block for the current batch
    const endBlock = Math.min(currentBlock + batchSize - 1, latestBlock);

    // Query events in the current batch
    const events = await myContract.queryFilter(
      "SimpleSwapDeployed",
      currentBlock,
      endBlock
    );
    for (let event of events) {
      const decodedData = await decodeTransactionInput(
        event.transactionHash,
        provider
      );
      if (decodedData) {
        decodedDataArray.push(decodedData);
      }
    }

    // Log completion of the current batch
    console.log(`Completed querying blocks ${currentBlock} to ${endBlock}`);

    // Wait for a specified delay before the next batch
    await new Promise((resolve) => setTimeout(resolve, delay));
  }

  // Save the decoded data array to a JSON file
  try {
    await fs.writeFile(
      "scripts/decodedData.json",
      JSON.stringify(decodedDataArray, replacer, 2)
    );
    console.log("Successfully wrote decoded data to decodedData.json");
  } catch (err) {
    console.error("Error writing file:", err);
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
