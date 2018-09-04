const promisify = (inner) => new Promise((resolve, reject) => inner((err, res) => err ? reject(err) : resolve(res)));
const getBalance = async (addr) => web3.utils.toBN(await web3.eth.getBalance(addr))

async function computeCost(receipt) {
  let { gasPrice } = await web3.eth.getTransaction(receipt.transactionHash)
  return web3.utils.toBN(gasPrice * receipt.gasUsed);
}

async function getTime() {
  return (await web3.eth.getBlock(await web3.eth.getBlockNumber())).timestamp
}

const isTestRPC = async () => (await web3.eth.getNodeInfo()).includes("TestRPC/v2")

const expectFail = async (promise) => {
  if(await isTestRPC()) {
    await promise.should.be.rejectedWith('revert');
  } else {
    throw 'figure out the current client behaviour'
  }
}

function matchLogs (logs, template) {
  if(logs.length != template.length) throw new Error('length does not match')
  for(let i = 0; i < logs.length; i++) {
    let log = logs[i]
    let temp = template[i]

    log.event.should.equal(temp.event)

    for(let arg in temp.args) {
      let v = temp.args[arg]
      if(typeof v === 'number') {
        v = web3.utils.toBN(v)
      }

      if(web3.utils.BN.isBN(v)) {
        log.args[arg].should.eq.BN(v)
      } else {
        log.args[arg].should.deep.equal(v)
      }
    }
  }
}

async function sign(signer, hash) {
  const sig = await web3.eth.sign(hash, signer);

  let r = sig.substr(0,66);
  let s = "0x" + sig.substr(66, 64);
  let v = parseInt(sig.substr(130, 2), 16) + 27

  return { r, s, v, sig }
}

const increaseTime = (sec) => {
  return promisify((cb) => {
    web3.currentProvider.send(
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
  nulladdress: '0x0000000000000000000000000000000000000000'
}
