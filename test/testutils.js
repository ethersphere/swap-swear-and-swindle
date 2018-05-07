const promisify = (inner) => new Promise((resolve, reject) => inner((err, res) => err ? reject(err) : resolve(res)));
const getBalance = (addr) => promisify((cb) => web3.eth.getBalance(addr, cb))
const getTransaction = (txHash) => promisify((cb) => web3.eth.getTransaction(txHash, cb))

async function computeCost(receipt) {
  let { gasPrice } = await getTransaction(receipt.transactionHash)
  return gasPrice.times(receipt.gasUsed);
}

let timeshift = web3.eth.getBlock(web3.eth.blockNumber).timestamp - Math.floor(Date.now() / 1000);

function getTime() {
  return Math.floor(Date.now() / 1000) + timeshift
}

const isTestRPC = () => web3.version.node.includes("TestRPC/v2")

const expectFail = async (promise) => {
  if(isTestRPC()) {
    await promise.should.be.rejectedWith('revert');
  } else {
    /* handle everything else */
    throw 'figure out the current client behaviour'
  }
}

function matchLogs (logs, template) {
  if(logs.length != template.length) throw new Error('length does not match')
  for(let i = 0; i < logs.length; i++) {
    let log = logs[i]
    let temp = template[i]

    for(let arg in temp.args) {
      if(typeof temp.args[arg] === 'number') {
        temp.args[arg] = web3.toBigNumber(temp.args[arg])
      }
    }

    log.args.should.deep.equal(temp.args)
    log.event.should.equal(temp.event)
  }
}

function sign(signer, hash) {
  const sig = web3.eth.sign(signer, hash);

  let r = sig.substr(0,66);
  let s = "0x" + sig.substr(66, 64);
  let v = parseInt(sig.substr(130, 2), 16) + 27

  return { r, s, v }
}

const increaseTime = (sec) => {
  timeshift += sec;
  return promisify((cb) => {
    web3.currentProvider.sendAsync(
      {
        jsonrpc: "2.0",
        method: "evm_increaseTime",
        params: [sec],
        id: 0
      }, cb)
  })
}

module.exports = {
  increaseTime,
  sign,
  matchLogs,
  expectFail,
  getTime,
  computeCost,
  getBalance,
  getTransaction,
  nulladdress: '0x0000000000000000000000000000000000000000'
}
