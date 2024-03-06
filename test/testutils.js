async function computeCost(receipt) {
  let { gasPrice } = await web3.eth.getTransaction(receipt.transactionHash);
  return web3.utils.toBN(gasPrice * receipt.gasUsed);
}

module.exports = {
  computeCost,
};
