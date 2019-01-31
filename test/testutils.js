const promisify = (inner) => new Promise((resolve, reject) => inner((err, res) => err ? reject(err) : resolve(res)));
const { time } = require('openzeppelin-test-helpers')

async function computeCost(receipt) {
  let { gasPrice } = await web3.eth.getTransaction(receipt.transactionHash)
  return web3.utils.toBN(gasPrice * receipt.gasUsed);
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
  matchLogs,
  computeCost,
  nulladdress: '0x0000000000000000000000000000000000000000'
}
